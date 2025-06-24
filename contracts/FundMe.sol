// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "./PriceConverter.sol";

error FundMe__NotOwner();
error FundMe__NotEnoughETH();
error FundMe__WithdrawFailed();
error FundMe__FundingClosed();
error FundMe__RefundFailed();
error FundMe__RefundPeriodOver();
error FundMe__NoFundsToRefund();

contract FundMe {
    using PriceConverter for uint256;

    uint256 public constant MINIMUM_USD = 5 * 1e18; // $5 in wei
    uint256 public constant FUNDING_PERIOD = 7 days; // Funding duration
    address private immutable i_owner;
    address[] private s_funders; // unique addresses that have funded
    mapping(address => uint256) private s_addressToAmountFunded;
    AggregatorV3Interface private s_priceFeed;
    uint256 public immutable s_startTime; // Funding start time

    modifier onlyOwner() {
        if (msg.sender != i_owner) revert FundMe__NotOwner();
        _;
    }

    modifier onlyWhileOpen() {
        if (block.timestamp > s_startTime + FUNDING_PERIOD) 
            revert FundMe__FundingClosed();
        _;
    }

    modifier onlyAfterClosed() {
        if (block.timestamp <= s_startTime + FUNDING_PERIOD) 
            revert FundMe__RefundPeriodOver();
        _;
    }

    constructor(address _priceFeed) {
        i_owner = msg.sender;
        s_priceFeed = AggregatorV3Interface(_priceFeed);
        s_startTime = block.timestamp;
    }

    function fund() public payable onlyWhileOpen {
        if (msg.value.getConversionRate(s_priceFeed) < MINIMUM_USD) {
            revert FundMe__NotEnoughETH();
        }
        
        // Only add new funders to the array
        if (s_addressToAmountFunded[msg.sender] == 0) {
            s_funders.push(msg.sender);
        }
        
        s_addressToAmountFunded[msg.sender] += msg.value;
    }

    // Fixed: Withdraw only after funding period ends
    function withdraw() public onlyOwner onlyAfterClosed {
        // Save funders array length before resetting
        uint256 fundersCount = s_funders.length;
        
        // Reset all funders' balances
        for (uint256 i = 0; i < fundersCount; i++) {
            address funder = s_funders[i];
            s_addressToAmountFunded[funder] = 0;
        }
        
        // Reset funders array
        s_funders = new address[](0);

        // Send all ETH to owner
        (bool success, ) = i_owner.call{value: address(this).balance}("");
        if (!success) revert FundMe__WithdrawFailed();
    }

    // Fixed: Only allow refunds while funding is open
    function refund() public onlyWhileOpen {
        uint256 amount = s_addressToAmountFunded[msg.sender];
        if (amount == 0) revert FundMe__NoFundsToRefund();
        
        // Reset before sending to prevent reentrancy
        s_addressToAmountFunded[msg.sender] = 0;
        
        // Remove from funders array
        for (uint256 i = 0; i < s_funders.length; i++) {
            if (s_funders[i] == msg.sender) {
                s_funders[i] = s_funders[s_funders.length - 1];
                s_funders.pop();
                break;
            }
        }

        (bool success, ) = msg.sender.call{value: amount}("");
        if (!success) {
            // Revert state if transfer fails
            s_addressToAmountFunded[msg.sender] = amount;
            revert FundMe__RefundFailed();
        }
    }

    function getFundingDeadline() public view returns (uint256) {
        return s_startTime + FUNDING_PERIOD;
    }
    
    function getRefundAvailable(address funder) public view returns (bool) {
        return block.timestamp <= s_startTime + FUNDING_PERIOD && 
            s_addressToAmountFunded[funder] > 0;
    }

    function getOwner() public view returns (address) {
        return i_owner;
    }

    // careful with querying large arrays
    function getFunder(uint256 index) public view returns (address) {
        return s_funders[index];
    }

    function getAddressToAmountFunded(address funder) public view returns (uint256) {
        return s_addressToAmountFunded[funder];
    }

    function getPriceFeed() public view returns (AggregatorV3Interface) {
        return s_priceFeed;
    }

    // unrecognized data, is processed as a funding contribution
    // (witout calldata)
    receive() external payable {
        fund();
    }
    // (with calldata but no matching function signature)
    fallback() external payable {
        fund();
    }
}