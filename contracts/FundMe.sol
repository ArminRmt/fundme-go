// collect funds (Ether)

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "./PriceConverter.sol";

error FundMe__NotOwner();
error FundMe__NotEnoughETH();
error FundMe__WithdrawFailed();

contract FundMe {
    using PriceConverter for uint256;

    uint256 public constant MINIMUM_USD = 5 * 1e18; // $5 in wei
    address private immutable i_owner;
    address[] private s_funders; // unique addresses that have funded
    mapping(address => uint256) private s_addressToAmountFunded;
    AggregatorV3Interface private s_priceFeed;

    modifier onlyOwner() {
        if (msg.sender != i_owner) revert FundMe__NotOwner();
        _; // execute the rest of the function called modifier
    }

    constructor(address _priceFeed) {
        i_owner = msg.sender;
        s_priceFeed = AggregatorV3Interface(_priceFeed); // state variable
    }

    // payable: allows the function to receive Ether
    function fund() public payable {
        if (msg.value.getConversionRate(s_priceFeed) < MINIMUM_USD) {
            revert FundMe__NotEnoughETH();
        }
        s_addressToAmountFunded[msg.sender] += msg.value;
        s_funders.push(msg.sender);
        // if (!s_funders[msg.sender]) {
        //     s_funders.push(msg.sender);
        // }

    }

    function withdraw() public onlyOwner {
        for (uint256 funderIndex = 0; funderIndex < s_funders.length; funderIndex++) {
            address funder = s_funders[funderIndex];
            s_addressToAmountFunded[funder] = 0;
        }
        s_funders = new address[](0);

        // Sending Ether to Owner
        (bool success, ) = i_owner.call{value: address(this).balance}("");
        if (!success) revert FundMe__WithdrawFailed();
    }

    function getOwner() public view returns (address) {
        return i_owner;
    }

    // Be careful with querying large arrays
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
