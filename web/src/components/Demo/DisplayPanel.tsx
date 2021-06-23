import { useState } from 'react';
import styles from './index.less';

import CanvasDraw from 'react-canvas-draw';
import classNames from './index.less';

const DisplayPanel: React.FC<PrintPanelProps> = (props) => {
    let loadableCanvas: any;

    const [color, updateColor] = useState<string>(props.color);
    const [width, updateWidth] = useState<number>(props.width);
    const [height, updateHeight] = useState<number>(props.height);
    const [brushRadius, updateBrushRadius] = useState<number>(props.brushRadius);
    const [lazyRadius, updateLazyRadius] = useState<number>(props.lazyRadius);
  
    return (
        <div>
            <button
            onClick={() => {
                loadableCanvas.loadSaveData(localStorage.getItem('savedDrawing'));
            }}
            >
            Load
            </button>
            <CanvasDraw
                disabled
                hideGrid
                canvasWidth={width}
                canvasHeight={height}
                ref={(canvasDraw: any) => (loadableCanvas = canvasDraw)}
                saveData={localStorage.getItem('savedDrawing')}
            />
        </div>
    );
  }

  export default DisplayPanel;