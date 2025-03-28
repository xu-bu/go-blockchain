// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {ERC721URIStorage, ERC721} from "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";


struct NFTListing {
  uint256 price;
  address seller;
}

contract NFTMarket is ERC721URIStorage {
  uint256 private _tokenIDs;
  mapping(uint256 => NFTListing) private _listings;

  // if tokenURI is not an empty string => an NFT was created
  // if price is not 0 => an NFT was listed
  // if price is 0 && tokenURI is an empty string => NFT was transferred (either bought, or the listing was canceled)
  event NFTTransfer(uint256 tokenID, address from, address to, string tokenURI, uint256 price);

  constructor() ERC721("Rock NFTs", "ANFT") {}

  function createNFT(string calldata tokenURI) public  {
      uint256 tokenID = _tokenIDs++;
      _safeMint(msg.sender, tokenID);
      _setTokenURI(tokenID, tokenURI);
      emit NFTTransfer(tokenID, address(0),msg.sender, tokenURI, 0);
  }

  function listNFT(uint256 tokenID, uint256 price) public {
    // require(price > 0, "NFTMarket: price must be greater than 0");
    transferFrom(msg.sender, address(this), tokenID);
    // _listings[tokenID] = NFTListing(price, msg.sender);
    emit NFTTransfer(tokenID, msg.sender, address(this), "", price);
  }

  function buyNFT(uint256 tokenID) public payable {
    //  NFTListing memory listing = _listings[tokenID];
    //  require(listing.price > 0, "NFTMarket: nft not listed for sale");
    //  require(msg.value == listing.price, "NFTMarket: incorrect price");
     ERC721(address(this)).transferFrom(address(this), msg.sender, tokenID);
    //  payable(listing.seller).transfer(listing.price.mul(95).div(100));
     emit NFTTransfer(tokenID, address(this), msg.sender, "", 0);
  }

  function withdrawFunds() public {
    uint256 balance =  address(this).balance;
    require(balance > 0, "NFTMarket: balance is zero");
    payable(msg.sender).transfer(balance); 
  }
}
