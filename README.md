1. make new address
2. trans with 0 value
3. messages
4. change address always new
5. nonce
6. tag

curl -H "Content-Type: application/json" -d '{"command": "getBalances", "addresses": ["Zoom"]}' http://localhost:3001


CREATE TABLE bundles (id bigint auto_increment, bundle varchar(81), ts datetime(6), 
                      last_index int, PRIMARY KEY(id), key(bundle), key(id, bundle, last_index, ts));

CREATE TABLE transactions (id bigint auto_increment, tx varchar(81), ts datetime(6), value bigint,
                           signature varchar(81),
                           current_index integer,
                           last_index integer,
                           bundle varchar(81),
                           trunk varchar(81),
                           branch varchar(81),
                           nonce varchar(81),
                           address varchar(81), PRIMARY KEY (id), 
                           KEY (address), KEY (tx), KEY (bundle));



"Trunk" and "Branch" are hashes of other transactions, namely the two transactions that were approved by the transaction you are currently looking at.

Bundle

    The 0th transaction is the spend transaction, where you transfer IOTAs to the recipient's address.

    The remaining transactions are inputs, where you transfer IOTAs from your addresses to fund the spend. 

    Each input transaction contains a signature fragment that proves that you own the private key 

    By default, signatures are so big that each one requires 2 transactions to store it. 
    This is configurable depending on the "security level" of the corresponding address — 
    it can be between 1 and 3. 
    If the sum of the inputs is greater than the spend, then an additional transaction is added to the end of the bundle that sends the change to an address that you own.


first you construct "raw" bundle
then you calculate bundle hash
then you sign that hash
and copy signatures into corresponding txs with correct offset

https://forum.iota.org/t/interesting-devs-conversation-on-bundles-transactions-hashing-of-pow-signature-in-iota/1399

does the bundle consists of tx ids?

there is no such structure as bundle, it's manifested only on meta-level
bundle is set of txs
so we can't say that bundle consists of tx ids, it consists of txs
which have ids/hashes (edited)


























ufw allow 8666

./simplecoin.life --gui no
./simplecoin.life --port 9999 --db zzz --peer 127.0.0.1:9376

1. s1 listens on 0.0.0.0
2. s2 connects to s1, s1 knows ip is 106.45.14.137


1. s1 starts
2. s2 starts, handshake 2<->1
3. s1 peers = [s2]
4. s2 peers = [s1]
5. s3 starts, handshake 3<->1
6. s1 peers = [s2,s3]
7. s3 peers = [s1,s2]
8. s4 starts, handshake 4<->1
9. s1 peers = [s3,s3,s4]
10. s4 peers = [s1,s2,s3]
11. s2 ping, 2<->1
12. s2 peers = [s1,s3,s4]
13. s3 ping, 3<->1!, 3<->2
14. s3 peers = [s1,s2,s4]

# For Fedora 26:

sudo yum install gtk3-devel
sudo yum install webkitgtk3-devel
sudo yum install mesa-libGLU-devel
sudo yum install libSM-devel
sudo yum install libnotify-devel
sudo yum install libjpeg-devel
sudo yum install libtiff-devel
sudo yum install gstreamer1-devel
sudo yum install gstream-devel

sudo yum install gstreamer1-plugins-base-devel

# simplecoin.life (SCL)

https://simplecoin.life

When you download and install the SCL client, you can generate your account number like:

Public Account Number: 47908853-30674977

Private Account Number: 47908853-32124409

You give out your Public Account Number to people as a way for them to send you money.
You keep your Private Account Number private.

A good first step after creating your numbers is to email both of them to yourself.

There is no way to reset your Private Account Number if you can't remember it.

If you title an email "my SCL account" and get that email into your cloud based sent-email folder, then
you could search for that email years later and find your private account number.

# transactions

a private key is a single unsigned 256 bit integer (32 bytes).

 A public key can be calculated from a private key, but not vice versa.

A public key can be used to determine if a signature is genuine (in other words, produced with the proper key) without requiring the private key to be divulged.

signature: A number that proves that a signing operation took place	

A signature is mathematically generated from a hash of something to be signed, plus a private key. The signature itself is two numbers known as r and s.

With the public key, a mathematical algorithm can be used on the signature to determine that it was originally produced from the hash and the private key


https://en.wikipedia.org/wiki/Elliptic_Curve_Digital_Signature_Algorithm

Bitcoin uses Elliptic Curve Digital Signature Algorithm (ECDSA) to sign transactions.
Signatures use DER encoding to pack the r and s components into a single byte stream 

Suppose Alice wants to send a signed message to Bob. Initially, they must agree on the curve parameters 
( CURVE , G , n)

we need G, a base point of prime order on the curve; 
n is the multiplicative order of the point G

Each input must have a cryptographic digital signature that unlocks the funds from the prior transaction. Only the person possessing the appropriate private key is able to create a satisfactory signature; this in effect ensures that funds can only be spent by their owners.

The sender generates a private key and public key. They then sign the message with the signature and send their public key, the signature and the message to the network (as the network is peer to peer each full node in the network validates each transaction) – The node or receiver then checks using the verification algorithm that the message has been signed by the sender, which can only be done by the holder of the private key to the public key that is sent.

GENESIS > andrew > hector > sue > bob

1. GENESIS
2. (hash of 1 + pubkey of andrew)
3. (hash of 2 + pubkey of hector)
4. (hash of 3 + pubkey of sue)
5. (hash of 4 + pubkey of bob)

* transactions must be publicly announced, and we need a system for participants to agree on a single history of the order in which they were received. The payee needs proof that at the time of each transaction it was the first receive

* chain of digital signatures. Each owner transfers the coin to the next by digitally signing a hash of the previous transaction and the public key of the next owner and adding these to the end of the coin.
* contain multiple inputs and outputs
* The timestamp proves that the data must have existed at the time, obviously, in order to get into the hash. Each timestamp includes the previous timestamp in its hash, forming a chain, with each additional timestamp reinforcing the ones before it.
*  incrementing a nonce in the block until a value is found that gives the block's hash the required zero bits. Once the CPU effort has been expended to make it satisfy the proof-of-work, the block cannot be changed without redoing the work. As later blocks are chained after it, the work to change the block would include redoing all the blocks after it.


https://en.bitcoin.it/wiki/Protocol_documentation#tx
TxID is just a SHA256 hash of binary transation data
https://medium.com/@lhartikk/a-blockchain-in-200-lines-of-code-963cc1cc0e54
https://github.com/lhartikk/naivechain/blob/master/main.js
