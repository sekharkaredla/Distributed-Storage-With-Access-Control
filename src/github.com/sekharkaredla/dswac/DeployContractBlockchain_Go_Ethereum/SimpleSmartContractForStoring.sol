pragma solidity >=0.4.22 <0.6.0;

contract SampleContract {

    string public message;

    constructor(string memory initialMessage) public {
        message = initialMessage;
    }

    function setMessage(string memory newMessage) public {
        message = newMessage;
    }
}