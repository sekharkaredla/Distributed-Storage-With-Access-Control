pragma solidity >=0.4.22 <0.6.0;

contract StoringDetails{

    bytes public JWT;
    function setJWT(bytes memory newJWT) public {
        JWT = newJWT;
    }
    function getJWT() public view returns (bytes memory){
        return JWT;
    }
}