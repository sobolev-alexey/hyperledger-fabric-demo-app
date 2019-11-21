import React, { Component } from 'react';
import { toast } from 'react-toastify';
import Notification from './Notification'; 
import './App.css';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      containerId: '',
      newHolderContainerId: '',
      newHolderName: '',
      newContainerId: '',
      containerDescription: '',
      longitude: '',
      latitude: '',
      holderName: '',
      allContainers: []
    };

    this.changeHolder = this.changeHolder.bind(this);
    this.createRecord = this.createRecord.bind(this);
    this.handleTextChange = this.handleTextChange.bind(this);
    this.queryContainer = this.queryContainer.bind(this);
    this.queryAllContainers = this.queryAllContainers.bind(this);
  }

  notifySuccess = message => toast.success(message);
  notifyError = message => toast.error(message);

  createRecord(event) {
    event.preventDefault();
    const { newContainerId, containerDescription, longitude, latitude, holderName } = this.state;
    if (newContainerId && containerDescription && longitude && latitude && holderName) {
      fetch('create', {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          key: newContainerId,
          holder: holderName,
          description: containerDescription,
          location: `${longitude}, ${latitude}`
        }),
      })
      .then(response => response.json())
      .then(data => {
        if (data.success && data.result) {
          this.notifySuccess('New record was created');
        } else {
          console.error(data.error);
          this.notifyError('Something went wrong');
        }
      });
    }
  }

  changeHolder(event) {
    event.preventDefault();
    const { allContainers, newHolderContainerId, newHolderName } = this.state;
    if (newHolderContainerId && newHolderName) {
      const container = allContainers.find(({ Key }) => Key === newHolderContainerId)
      if (container && container.Record.holder === 'Retailer') {
        console.error('Container arrived to retailer. No further change possible');
        this.notifyError('Container arrived to retailer. No further change possible');
        return
      }

      fetch('change', {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          id: newHolderContainerId,
          holder: newHolderName
        }),
      })
      .then(response => response.json())
      .then(data => {
        if (data.success && data.result) {
          this.notifySuccess('Holder was changed');
        } else {
          console.error(data.error);
          this.notifyError('Something went wrong');
        }
      });
    }
  }

  handleTextChange = event => {
    this.setState({ [event.target.id]: event.target.value });
  };

  queryContainer(event) {
    event.preventDefault();
    const { containerId } = this.state;
    if (containerId) {
      fetch(`get/${encodeURIComponent(containerId)}`)
        .then(response => response.json())
        .then(data => {
          if (data.success && data.result) {
            const result = JSON.parse(data.result)

            this.setState({ 
              allContainers: [{ Key: containerId, Record: result.container }]
            });
            console.log("Wallet Address", result.wallet)
            console.log("MAM Root", result.mamstate.root)
            console.log("MAM payload: ")
            console.log(result.messages)
            console.log("======================")
          } else {
            console.error(data.error);
          }
      });
    }
  }

  queryAllContainers(event) {
    event.preventDefault();
    fetch('get_all')
      .then(response => response.json())
      .then(data => {
        if (data.success && data.result) {
          this.setState({ allContainers: data.result });
        } else {
          console.error(data.error);
        }
      });
  }

  render() {
    return (
      <div className="App">
        <header>
          <div id="left_header">Hyperledger Fabric Demo Application</div>
          <i id="right_header">Example Blockchain Application for Hyperledger Fabric</i>
        </header>
        <div className="queryContainer">
          <form onSubmit={this.queryContainer}>
            <label>Query a Specific Container</label><br />
            Enter a container ID: <br />
            <input
              id="containerId"
              type="number"
              placeholder="Ex: 3"
              value={this.state.containerId}
              onChange={this.handleTextChange}
            />
            <br />
            <button type="submit" className="btn btn-primary">Query Container Record</button>
          </form>
        </div>
        <br />
        <br />

        <div className="queryAllContainers">
          <div className="form-group">
            <label>Query All Containers</label><br />
            <button type="button" className="btn btn-primary" onClick={this.queryAllContainers}>Query All Containers</button>
          </div>
 
          {
            this.state.allContainers.length ? (
              <table id="all_containers" className="table" align="center">
                <thead>
                  <tr>
                    <th>ID</th>
                    <th>Timestamp</th>
                    <th>Holder</th>
                    <th>Container Location (Longitude, Latitude)</th>
                    <th>Container Description</th>
                  </tr>
                </thead>
                <tbody>
                  {
                    this.state.allContainers
                    .sort((a, b) => parseFloat(a.Key) - parseFloat(b.Key))
                    .map(container => (
                      <tr key={container.Key}>
                        <td>{container.Key}</td>
                        <td>{container.Record.timestamp}</td>
                        <td>{container.Record.holder}</td>
                        <td>{container.Record.location}</td>
                        <td>{container.Record.description}</td>
                      </tr>
                    ))
                  }
                </tbody>
              </table>
            ) : null
          }
        </div>

        <br />
        <br />

        <div className="createRecord">
          <form onSubmit={this.createRecord}>
            <label>Create Container Record</label>
            <br />
            Enter container id:
            <input
              className="form-control" 
              id="newContainerId"
              name="newContainerId" 
              type="text" 
              placeholder="Ex: 11" 
              value={this.state.newContainerId}
              onChange={this.handleTextChange}
            />
            Enter container description: 
            <input 
              className="form-control" 
              id="containerDescription"
              name="containerDescription" 
              type="text" 
              placeholder="Ex: 0239L" 
              value={this.state.containerDescription}
              onChange={this.handleTextChange}
            />
            Enter longitude: 
            <input 
              className="form-control" 
              id="longitude"
              name="longitude" 
              type="number" 
              placeholder="Ex: 28.012"
              value={this.state.longitude}
              onChange={this.handleTextChange}
            /> 
            Enter latitude: 
            <input 
              className="form-control" 
              id="latitude"
              name="latitude" 
              type="number" 
              placeholder="Ex: 150.405"
              value={this.state.latitude}
              onChange={this.handleTextChange}
            />
            Enter name of holder: 
            <input 
              className="form-control" 
              id="holderName"
              name="holderName" 
              type="text" 
              placeholder="Ex: Hansel" 
              value={this.state.holderName}
              onChange={this.handleTextChange}
            />
            <button type="submit" className="btn btn-primary">Create record</button>
          </form>
        </div>

        <br />
        <br />

        <div className="changeContainerHolder">
          <form onSubmit={this.changeHolder}>
            <label>Change Container Holder</label><br />
             Enter a container ID:
            <input
              className="form-control"
              id="newHolderContainerId"
              name="newHolderContainerId" 
              placeholder="Ex: 1" 
              type="number"
              value={this.state.newHolderContainerId}
              onChange={this.handleTextChange}
            />
            Enter name of new holder:
            <input
              className="form-control"
              id="newHolderName"
              name="newHolderName"
              placeholder="Ex: Barry"
              type="text"
              value={this.state.newHolderName}
              onChange={this.handleTextChange}
            />
            <button type="submit" className="btn btn-primary">Change</button>
          </form>
        </div>
        
        <br /><br /><br />
        <Notification />
      </div>
    );
  }
}

export default App;
