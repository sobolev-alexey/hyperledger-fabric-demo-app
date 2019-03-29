import React, { Component } from 'react';
import { toast } from 'react-toastify';
import Notification from './Notification'; 
import './App.css';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      tunaId: '',
      newHolderTunaId: '',
      newHolderName: '',
      newTunaId: '',
      vesselName: '',
      longitude: '',
      latitude: '',
      holderName: '',
      allTuna: []
    };

    this.changeHolder = this.changeHolder.bind(this);
    this.createRecord = this.createRecord.bind(this);
    this.handleTextChange = this.handleTextChange.bind(this);
    this.queryTuna = this.queryTuna.bind(this);
    this.queryAllTuna = this.queryAllTuna.bind(this);
  }

  notifySuccess = message => toast.success(message);
  notifyError = message => toast.error(message);

  createRecord(event) {
    event.preventDefault();
    const { newTunaId, vesselName, longitude, latitude, holderName } = this.state;
    if (newTunaId && vesselName && longitude && latitude && holderName) {
      fetch('create', {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          key: newTunaId,
          holder: holderName,
          vessel: vesselName,
          location: `${longitude}, ${latitude}`,
          timestamp: Date.now().toString()
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
    const { newHolderTunaId, newHolderName } = this.state;
    if (newHolderTunaId && newHolderName) {
      fetch('change', {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          id: newHolderTunaId,
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

  queryTuna(event) {
    event.preventDefault();
    const { tunaId } = this.state;
    if (tunaId) {
      fetch(`get/${encodeURIComponent(tunaId)}`)
        .then(response => response.json())
        .then(data => {
          if (data.success && data.result) {
            this.setState({ 
              allTuna: [{ Key: tunaId, Record: JSON.parse(data.result) }]
            });
          } else {
            console.error(data.error);
          }
      });
    }
  }

  queryAllTuna(event) {
    event.preventDefault();
    fetch('get_all')
      .then(response => response.json())
      .then(data => {
        if (data.success && data.result) {
          this.setState({ allTuna: data.result });
        } else {
          console.error(data.error);
        }
      });
  }

  render() {
    return (
      <div className="App">
        <header>
          <div id="left_header">Hyperledger Fabric Tuna Application</div>
          <i id="right_header">Example Blockchain Application for Hyperledger Fabric</i>
        </header>
        <div className="queryTuna">
          <form onSubmit={this.queryTuna}>
            <label>Query a Specific Tuna Catch</label><br />
            Enter a catch number: <br />
            <input
              id="tunaId"
              type="number"
              placeholder="Ex: 3"
              value={this.state.tunaId}
              onChange={this.handleTextChange}
            />
            <br />
            <button type="submit" className="btn btn-primary">Query Tuna Record</button>
          </form>
        </div>
        <br />
        <br />

        <div className="queryAllTuna">
          <div className="form-group">
            <label>Query All Tuna Catches</label><br />
            <button type="button" className="btn btn-primary" onClick={this.queryAllTuna}>Query All Tuna</button>
          </div>
 
          {
            this.state.allTuna.length ? (
              <table id="all_tuna" className="table" align="center">
                <thead>
                  <tr>
                    <th>ID</th>
                    <th>Timestamp</th>
                    <th>Holder</th>
                    <th>Catch Location (Longitude, Latitude)</th>
                    <th>Vessel</th>
                  </tr>
                </thead>
                <tbody>
                  {
                    this.state.allTuna
                    .sort((a, b) => parseFloat(a.Key) - parseFloat(b.Key))
                    .map(tuna => (
                      <tr key={tuna.Key}>
                        <td>{tuna.Key}</td>
                        <td>{tuna.Record.timestamp}</td>
                        <td>{tuna.Record.holder}</td>
                        <td>{tuna.Record.location}</td>
                        <td>{tuna.Record.vessel}</td>
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
            <label>Create Tuna Record</label>
            {/* <h5 style="color:green;margin-bottom:2%" id="success_create">Success! Tx ID: {create_tuna}</h5> */}
            <br />
            Enter catch id:
            <input
              className="form-control" 
              id="newTunaId"
              name="newTunaId" 
              type="text" 
              placeholder="Ex: 11" 
              value={this.state.newTunaId}
              onChange={this.handleTextChange}
            />
            Enter name of vessel: 
            <input 
              className="form-control" 
              id="vesselName"
              name="vesselName" 
              type="text" 
              placeholder="Ex: 0239L" 
              value={this.state.vesselName}
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

        <div className="changeTunaHolder">
          <form onSubmit={this.changeHolder}>
            <label>Change Tuna Holder</label><br />
            {/* <h5 style="color:green;margin-bottom:2%" id="success_holder">Success! Tx ID: {change_holder}</h5>
            <h5 style="color:red;margin-bottom:2%" id="error_holder">Error: Please enter a valid Tuna Id</h5> */}
            Enter a catch id between 1 and 10:
            <input
              className="form-control"
              id="newHolderTunaId"
              name="newHolderTunaId" 
              placeholder="Ex: 1" 
              type="number"
              value={this.state.newHolderTunaId}
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
