import React, { Component } from 'react';
import styles from './print.less';

import CanvasDraw from "react-canvas-draw"
import classNames from "./print.less";

class Demo extends Component {
    state = {
      color: "#ffc600",
      width: 1000,
      height: 1000,
      brushRadius: 10,
      lazyRadius: 12
    };
    // componentDidMount() {
    //   // let's change the color randomly every 2 seconds. fun!
    //   window.setInterval(() => {
    //     this.setState({
    //       color: "#" + Math.floor(Math.random() * 16777215).toString(16)
    //     });
    //   }, 2000);
    // }
    render() {
      return (
        <div>
          <button
            onClick={() => {
              this.setState({
                color: "#ff0000"
              })
            }}
          >
            Red
          </button>
  
          <button
            onClick={() => {
              this.setState({
                color: "#00ff00"
              })
            }}
          >
            Green
          </button>
  
          <button
            onClick={() => {
              this.setState({
                color: "#0000ff"
              })
            }}
          >
            Blue
          </button>
            
          <div className={classNames.tools}>
            <button
              onClick={() => {
                localStorage.setItem(
                  "savedDrawing",
                  this.saveableCanvas.getSaveData()
                );
              }}
            >
              Save
            </button>
            <button
              onClick={() => {
                this.saveableCanvas.clear();
              }}
            >
              Clear
            </button>
            <button
              onClick={() => {
                this.saveableCanvas.undo();
              }}
            >
              Reset
            </button>
            <div>
              <label>Width:</label>
              <input
                type="number"
                value={this.state.width}
                onChange={e =>
                  this.setState({ width: parseInt(e.target.value, 10) })
                }
              />
            </div>
            <div>
              <label>Height:</label>
              <input
                type="number"
                value={this.state.height}
                onChange={e =>
                  this.setState({ height: parseInt(e.target.value, 10) })
                }
              />
            </div>
            <div>
              <label>Brush-Radius:</label>
              <input
                type="number"
                value={this.state.brushRadius}
                onChange={e =>
                  this.setState({ brushRadius: parseInt(e.target.value, 10) })
                }
              />
            </div>
            <div>
              <label>Lazy-Radius:</label>
              <input
                type="number"
                value={this.state.lazyRadius}
                onChange={e =>
                  this.setState({ lazyRadius: parseInt(e.target.value, 10) })
                }
              />
            </div>
          </div>
          <CanvasDraw
            ref={canvasDraw => (this.saveableCanvas = canvasDraw)}
            brushColor={this.state.color}
            brushRadius={this.state.brushRadius}
            lazyRadius={this.state.lazyRadius}
            canvasWidth={this.state.width}
            canvasHeight={this.state.height}
            onChange={() => console.log("onChange")}
          />
          <button
          onClick={() => {
              this.loadableCanvas.loadSaveData(
                localStorage.getItem("savedDrawing")
              );
            }}
          >
            Load
          </button>
          <CanvasDraw
            disabled
            hideGrid
            canvasWidth={this.state.width}
            canvasHeight={this.state.height}
            ref={canvasDraw => (this.loadableCanvas = canvasDraw)}
            saveData={localStorage.getItem("savedDrawing")}
          />
        </div>
      );
    }
  }

const Print: React.FC = () => {
    return (
        <Demo />
    )
}

export default Print;