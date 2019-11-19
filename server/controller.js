//SPDX-License-Identifier: Apache-2.0

/*
  This code is based on code written by the Hyperledger Fabric community.
  Original code can be found here: https://github.com/hyperledger/fabric-samples/blob/release/fabcar/query.js
  and https://github.com/hyperledger/fabric-samples/blob/release/fabcar/invoke.js
 */

const Fabric_Client = require('fabric-client');
const path          = require('path');
const util          = require('util');
const os            = require('os');

module.exports = {
	get_all_containers: (req, res) => {
		console.log("getting all containers from database: ");

		const fabric_client = new Fabric_Client();

		// setup the fabric network
		const channel = fabric_client.newChannel('mychannel');
		const peer = fabric_client.newPeer('grpc://localhost:7051');
		channel.addPeer(peer);

		const store_path = path.join(os.homedir(), '.hfc-key-store');
		console.log('Store path:'+store_path);
		const tx_id = null;

		// create the key value store as defined in the fabric-client/config/default.json 'key-value-store' setting
		Fabric_Client.newDefaultKeyValueStore({ path: store_path })
		.then(state_store => {
		    // assign the store to the fabric client
		    fabric_client.setStateStore(state_store);
		    const crypto_suite = Fabric_Client.newCryptoSuite();
		    // use the same location for the state store (where the users' certificate are kept)
		    // and the crypto store (where the users' keys are kept)
		    const crypto_store = Fabric_Client.newCryptoKeyStore({ path: store_path });
		    crypto_suite.setCryptoKeyStore(crypto_store);
		    fabric_client.setCryptoSuite(crypto_suite);

		    // get the enrolled user from persistence, this user will sign all requests
		    return fabric_client.getUserContext('adminUser', true);
		}).then(user_from_store => {
		    if (user_from_store && user_from_store.isEnrolled()) {
		        console.log('Successfully loaded adminUser from persistence');
		    } else {
		        throw new Error('Failed to get adminUser.... run registerUser.js');
		    }

		    // queryAllContainers - requires no arguments , ex: args: [''],
		    const request = {
		        chaincodeId: 'fabric-demo-app',
		        txId: tx_id,
		        fcn: 'queryAllContainers',
		        args: ['']
		    };

		    // send the query proposal to the peer
		    return channel.queryByChaincode(request);
		}).then(query_responses => {
		    console.log("Query has completed, checking results");
		    // query_responses could have more than one  results if there multiple peers were used as targets
		    if (query_responses && query_responses.length === 1) {
		        if (query_responses[0] instanceof Error) {
					console.error("error from query = ", query_responses[0]);
					res.send({ success: false, error: `Could not locate containers. ${query_responses[0]}` });
		        } else {
		            console.log("Response is", query_responses[0].toString());
		            res.send({ success: true, result: JSON.parse(query_responses[0].toString()) });
		        }
		    } else {
				console.log("No payloads were returned from query");
				res.send({ success: false, error: 'No payloads were returned from query' });
		    }
		}).catch(err => {
			console.error('Failed to query successfully :: ' + err);
			res.send({ success: false, error: `Failed to query all containers. ${err}` });
		});
	},
	get_container: (req, res) => {
		const fabric_client = new Fabric_Client();
		const key = req.params.id;
		
		// setup the fabric network
		const channel = fabric_client.newChannel('mychannel');
		const peer = fabric_client.newPeer('grpc://localhost:7051');
		channel.addPeer(peer);

		const store_path = path.join(os.homedir(), '.hfc-key-store');
		console.log('Store path:' + store_path);
		const tx_id = null;

		// create the key value store as defined in the fabric-client/config/default.json 'key-value-store' setting
		Fabric_Client.newDefaultKeyValueStore({ path: store_path })
		.then(state_store => {
		    // assign the store to the fabric client
		    fabric_client.setStateStore(state_store);
		    const crypto_suite = Fabric_Client.newCryptoSuite();
		    // use the same location for the state store (where the users' certificate are kept)
		    // and the crypto store (where the users' keys are kept)
		    const crypto_store = Fabric_Client.newCryptoKeyStore({path: store_path});
		    crypto_suite.setCryptoKeyStore(crypto_store);
		    fabric_client.setCryptoSuite(crypto_suite);

		    // get the enrolled user from persistence, this user will sign all requests
		    return fabric_client.getUserContext('adminUser', true);
		}).then(user_from_store => {
		    if (user_from_store && user_from_store.isEnrolled()) {
		        console.log('Successfully loaded adminUser from persistence');
		    } else {
		        throw new Error('Failed to get adminUser.... run registerUser.js');
		    }

		    // queryContainer - requires 1 argument, ex: args: ['4'],
		    const request = {
		        chaincodeId: 'fabric-demo-app',
		        txId: tx_id,
		        fcn: 'queryContainer',
		        args: [key]
		    };

		    // send the query proposal to the peer
		    return channel.queryByChaincode(request);
		}).then(query_responses => {
		    console.log("Query has completed, checking results");
		    // query_responses could have more than one  results if there multiple peers were used as targets
		    if (query_responses && query_responses.length === 1) {
		        if (query_responses[0] instanceof Error) {
		            console.error("error from query = ", query_responses[0]);
		            res.send({ success: false, error: `Could not locate container. ${query_responses[0]}` })

		        } else {
		            console.log("Response is", query_responses[0].toString());
		            res.send({ success: true, result: query_responses[0].toString() });
		        }
		    } else {
		        console.log("No payloads were returned from query");
		        res.send({ success: false, error: 'Could not locate container. No payloads were returned from query' });
		    }
		}).catch(err => {
		    console.error('Failed to query successfully :: ' + err);
		    res.send({ success: false, error: `Could not locate container. ${err}` });
		});
	},
	create_record: (req, res) => {
		console.log("submit recording of a container: ");

		const { key, location, description, holder } = req.body;
		const fabric_client = new Fabric_Client();

		// setup the fabric network
		const channel = fabric_client.newChannel('mychannel');
		const peer = fabric_client.newPeer('grpc://localhost:7051');
		channel.addPeer(peer);
		const order = fabric_client.newOrderer('grpc://localhost:7050')
		channel.addOrderer(order);

		const store_path = path.join(os.homedir(), '.hfc-key-store');
		console.log('Store path:'+store_path);
		let tx_id = null;

		// create the key value store as defined in the fabric-client/config/default.json 'key-value-store' setting
		Fabric_Client.newDefaultKeyValueStore({ path: store_path })
		.then(state_store => {
		    // assign the store to the fabric client
		    fabric_client.setStateStore(state_store);
		    const crypto_suite = Fabric_Client.newCryptoSuite();
		    // use the same location for the state store (where the users' certificate are kept)
		    // and the crypto store (where the users' keys are kept)
		    const crypto_store = Fabric_Client.newCryptoKeyStore({ path: store_path });
		    crypto_suite.setCryptoKeyStore(crypto_store);
		    fabric_client.setCryptoSuite(crypto_suite);

		    // get the enrolled user from persistence, this user will sign all requests
		    return fabric_client.getUserContext('adminUser', true);
		}).then(user_from_store => {
		    if (user_from_store && user_from_store.isEnrolled()) {
		        console.log('Successfully loaded adminUser from persistence');
		    } else {
		        throw new Error('Failed to get adminUser.... run registerUser.js');
		    }

		    // get a transaction id object based on the current user assigned to fabric client
		    tx_id = fabric_client.newTransactionID();
		    console.log("Assigning transaction_id: ", tx_id._transaction_id);

		    // recordContainer - requires 4 args, ID, description, location, holder - ex: args: ['10', 'Apples', '-12.021, 28.012', 'Hansel'],
			// send proposal to endorser
			// https://fabric-sdk-node.github.io/global.html#ChaincodeInvokeRequest
		    const request = {
		        //targets : --- letting this default to the peers assigned to the channel
		        chaincodeId: 'fabric-demo-app',
		        fcn: 'recordContainer',
		        args: [key, description, location, holder],
		        chainId: 'mychannel',
		        txId: tx_id
		    };

			// send the transaction proposal to the peers
			// https://fabric-sdk-node.github.io/Channel.html#sendTransactionProposal__anchor
		    return channel.sendTransactionProposal(request);
		}).then(results => {
		    const proposalResponses = results[0];
		    const proposal = results[1];
		    let isProposalGood = false;
		    if (proposalResponses && proposalResponses[0].response &&
		        proposalResponses[0].response.status === 200) {
		            isProposalGood = true;
		            console.log('Transaction proposal was good');
		        } else {
		            console.error('Transaction proposal was bad');
		        }
		    if (isProposalGood) {
		        console.log(util.format(
		            'Successfully sent Proposal and received ProposalResponse: Status - %s, message - "%s"',
		            proposalResponses[0].response.status, proposalResponses[0].response.message));

		        // build up the request for the orderer to have the transaction committed
		        const request = {
		            proposalResponses,
		            proposal
		        };

		        // set the transaction listener and set a timeout of 30 sec
		        // if the transaction did not get committed within the timeout period,
		        // report a TIMEOUT status
		        const transaction_id_string = tx_id.getTransactionID(); //Get the transaction ID string to be used by the event processing
		        const promises = [];

		        const sendPromise = channel.sendTransaction(request);
		        promises.push(sendPromise); //we want the send transaction first, so that we know where to check status

		        // get an eventhub once the fabric client has a user assigned. The user
		        // is required because the event registration must be signed
				let event_hub = channel.newChannelEventHub(peer);

		        // using resolve the promise so that result status may be processed
		        // under the then clause rather than having the catch clause process
		        // the status
		        let txPromise = new Promise((resolve, reject) => {
		            let handle = setTimeout(() => {
		                event_hub.disconnect();
		                resolve({ event_status: 'TIMEOUT' }); // we could use reject(new Error('Transaction did not complete within 30 seconds'));
		            }, 3000);
		            event_hub.connect();
		            event_hub.registerTxEvent(transaction_id_string, (tx, code) => {
		                // this is the callback for transaction event status
		                // first some clean up of event listener
		                clearTimeout(handle);
		                event_hub.unregisterTxEvent(transaction_id_string);
		                event_hub.disconnect();

		                // now let the application know what happened
		                const return_status = { event_status: code, tx_id: transaction_id_string };
		                if (code !== 'VALID') {
		                    console.error('The transaction was invalid, code = ' + code);
		                    resolve(return_status); // we could use reject(new Error('Problem with the tranaction, event status ::'+code));
		                } else {
		                    console.log('The transaction has been committed on peer ' + event_hub._peer._endpoint.addr);
		                    resolve(return_status);
		                }
		            }, err => {
		                //this is the callback if something goes wrong with the event registration or processing
		                reject(new Error('There was a problem with the eventhub ::' + err));
		            });
		        });
		        promises.push(txPromise);

		        return Promise.all(promises);
		    } else {
		        console.error('Failed to send Proposal or receive valid response. Response null or status is not 200. exiting...');
		        throw new Error('Failed to send Proposal or receive valid response. Response null or status is not 200. exiting...');
		    }
		}).then(results => {
		    console.log('Send transaction promise and event listener promise have completed');
		    // check the results in the order the promises were added to the promise all list
			if (results && results[0] && results[0].status) {
				if (results[0].status === 'SUCCESS') {
					console.log('Successfully sent transaction to the orderer.');

					if (results[1] && results[1].event_status) {
						if (results[1].event_status === 'VALID') {
							console.log('Successfully committed the change to the ledger by the peer');
							res.send({ success: true, result: tx_id.getTransactionID() });
						} else {
							console.log('Transaction failed to be committed to the ledger due to ::' + results[1].event_status);
							res.send({ success: false, error: `Transaction failed to be committed to the ledger due to ${results[1].event_status}` });
						}
					}	
				} else {
					console.error('Failed to order the transaction. Error code: ' + results[0].status);
					res.send({ success: false, error: `Failed to order the transaction. Error code: ${results[0].status}` });
				}
			}
		}).catch(err => {
			console.error('Failed to invoke successfully :: ' + err);
			res.send({ success: false, error: `Failed to invoke create_record. ${err}` });
		});
	},
	change_holder: (req, res) => {
		console.log("changing holder of a container:", req.params);

		const containerId = req.body.id;
		const holder = req.body.holder;
		const fabric_client = new Fabric_Client();

		// setup the fabric network
		const channel = fabric_client.newChannel('mychannel');
		const peer = fabric_client.newPeer('grpc://localhost:7051');
		channel.addPeer(peer);
		const order = fabric_client.newOrderer('grpc://localhost:7050')
		channel.addOrderer(order);

		// var member_user = null;
		const store_path = path.join(os.homedir(), '.hfc-key-store');
		console.log('Store path:'+store_path);
		let tx_id = null;

		// create the key value store as defined in the fabric-client/config/default.json 'key-value-store' setting
		Fabric_Client.newDefaultKeyValueStore({ path: store_path })
		.then((state_store) => {
		    // assign the store to the fabric client
		    fabric_client.setStateStore(state_store);
		    const crypto_suite = Fabric_Client.newCryptoSuite();
		    // use the same location for the state store (where the users' certificate are kept)
		    // and the crypto store (where the users' keys are kept)
		    const crypto_store = Fabric_Client.newCryptoKeyStore({ path: store_path });
		    crypto_suite.setCryptoKeyStore(crypto_store);
		    fabric_client.setCryptoSuite(crypto_suite);

		    // get the enrolled user from persistence, this user will sign all requests
		    return fabric_client.getUserContext('adminUser', true);
		}).then(user_from_store => {
		    if (user_from_store && user_from_store.isEnrolled()) {
		        console.log('Successfully loaded adminUser from persistence');
		    } else {
		        throw new Error('Failed to get adminUser.... run registerUser.js');
		    }

		    // get a transaction id object based on the current user assigned to fabric client
		    tx_id = fabric_client.newTransactionID();
		    console.log("Assigning transaction_id: ", tx_id._transaction_id);

		    // changeContainerHolder - requires 2 args , ex: args: ['1', 'Barry'],
		    // send proposal to endorser
		    const request = {
		        //targets : --- letting this default to the peers assigned to the channel
		        chaincodeId: 'fabric-demo-app',
		        fcn: 'changeContainerHolder',
		        args: [containerId, holder],
		        chainId: 'mychannel',
		        txId: tx_id
		    };

			// send the transaction proposal to the peers
		    return channel.sendTransactionProposal(request);
		}).then(results => {
		    const proposalResponses = results[0];
		    const proposal = results[1];
		    let isProposalGood = false;
			if (proposalResponses && 
				proposalResponses[0].response &&
		        proposalResponses[0].response.status === 200) {
				isProposalGood = true;
				console.log('Transaction proposal was good');
			} else {
				console.error('Transaction proposal was bad');
			}
		    if (isProposalGood) {
		        console.log(util.format(
		            'Successfully sent Proposal and received ProposalResponse: Status - %s, message - "%s"',
		            proposalResponses[0].response.status, proposalResponses[0].response.message));

		        // build up the request for the orderer to have the transaction committed
		        const request = {
		            proposalResponses: proposalResponses,
		            proposal: proposal
		        };

		        // set the transaction listener and set a timeout of 30 sec
		        // if the transaction did not get committed within the timeout period,
		        // report a TIMEOUT status
		        const transaction_id_string = tx_id.getTransactionID(); //Get the transaction ID string to be used by the event processing
				const promises = [];

				// https://fabric-sdk-node.github.io/Channel.html#sendTransactionProposal__anchor
		        const sendPromise = channel.sendTransaction(request);
		        promises.push(sendPromise); //we want the send transaction first, so that we know where to check status

		        // get an eventhub once the fabric client has a user assigned. The user
		        // is required because the event registration must be signed
		        let event_hub = channel.newChannelEventHub(peer);

		        // using resolve the promise so that result status may be processed
		        // under the then clause rather than having the catch clause process
		        // the status
		        let txPromise = new Promise((resolve, reject) => {
		            const handle = setTimeout(() => {
		                event_hub.disconnect();
		                resolve({event_status : 'TIMEOUT'}); // we could use reject(new Error('Transaction did not complete within 30 seconds'));
					}, 30000);
		            event_hub.connect();
		            event_hub.registerTxEvent(transaction_id_string, (tx, code) => {
		                // this is the callback for transaction event status
		                // first some clean up of event listener
						clearTimeout(handle);
		                event_hub.unregisterTxEvent(transaction_id_string);
						event_hub.disconnect();

		                // now let the application know what happened
		                const return_status = { event_status: code, tx_id: transaction_id_string };
						if (code !== 'VALID') {
		                    console.error('The transaction was invalid, code = ' + code);
		                    resolve(return_status); // we could use reject(new Error('Problem with the transaction, event status ::'+code));
		                } else {
		                    console.log('The transaction has been committed on peer ' + event_hub._peer._endpoint.addr);
		                    resolve(return_status);
						}
		            }, err => {
		                // this is the callback if something goes wrong with the event registration or processing
		                reject(new Error('There was a problem with the eventhub ::' + err));
		            });
				});
				promises.push(txPromise);
		        return Promise.all(promises);
		    } else {
		        console.error('Failed to send Proposal or receive valid response. Response null or status is not 200. exiting...');
		        throw new Error('Failed to send Proposal or receive valid response. Response null or status is not 200. exiting...');
		    }
		}).then(results => {
		    console.log('Send transaction promise and event listener promise have completed');
		    // check the results in the order the promises were added to the promise all list
			if (results && results[0] && results[0].status) {
				if (results[0].status === 'SUCCESS') {
					console.log('Successfully sent transaction to the orderer.');

					if (results[1] && results[1].event_status) {
						if (results[1].event_status === 'VALID') {
							console.log('Successfully committed the change to the ledger by the peer');
							res.send({ success: true, result: tx_id.getTransactionID() });
						} else {
							console.log('Transaction failed to be committed to the ledger due to ::' + results[1].event_status);
							res.send({ success: false, error: `Transaction failed to be committed to the ledger due to ${results[1].event_status}` });
						}
					}
				} else {
					console.error('Failed to order the transaction. Error code: ' + results[0].status);
					res.send({ success: false, error: `Failed to order the transaction. Error code: ${results[0].status}` });
				}
			}
		}).catch(err => {
			console.error('Failed to invoke successfully :: ' + err);
			res.send({ success: false, error: `Failed to invoke change_holder. ${err}` });
		});
	}
}
