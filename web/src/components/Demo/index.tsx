import { useState, useRef } from 'react';
import styles from './index.less';

import CanvasDraw, { CanvasDrawProps } from 'react-canvas-draw';
import classNames from './index.less';

export type PrintPanelProps = {
  color: string;
  width: number;
  height: number;
  brushRadius: number;
  lazyRadius: number;
};

const PrintPanel: React.FC<PrintPanelProps> = (props) => {
  // const saveableCanvas = useRef<CanvasDraw>();
  let saveableCanvas: any;

  const [color, updateColor] = useState<string>(props.color);
  const [width, updateWidth] = useState<number>(props.width);
  const [height, updateHeight] = useState<number>(props.height);
  const [brushRadius, updateBrushRadius] = useState<number>(props.brushRadius);
  const [lazyRadius, updateLazyRadius] = useState<number>(props.lazyRadius);

  return (
    <div>
      <button
        onClick={() => (
          updateColor('#ff0000')
        )}
      >
        Red
      </button>

      <button
        onClick={() => (
          updateColor('#00ff00')
        )}
      >
        Green
      </button>

      <button
        onClick={() => (
          updateColor('#0000ff')
        )}
      >
        Blue
      </button>

      <div className={classNames.tools}>
        <button
          onClick={() => {
            localStorage.setItem('savedDrawing', saveableCanvas.getSaveData());
          }}
        >
          Save
        </button>
        <button
          onClick={() => {
            saveableCanvas.clear();
          }}
        >
          Clear
        </button>
        <button
          onClick={() => {
            saveableCanvas.undo();
          }}
        >
          Reset
        </button>
        <div>
          <label>Width:</label>
          <input
            type="number"
            value={width}
            onChange={(e) => ( updateWidth(parseInt(e.target.value, 10)) )}
          />
        </div>
        <div>
          <label>Height:</label>
          <input
            type="number"
            value={height}
            onChange={(e) => ( updateHeight(parseInt(e.target.value, 10)) )}
          />
        </div>
        <div>
          <label>Brush-Radius:</label>
          <input
            type="number"
            value={brushRadius}
            onChange={(e) => ( updateBrushRadius(parseInt(e.target.value, 10)) )}
          />
        </div>
        <div>
          <label>Lazy-Radius:</label>
          <input
            type="number"
            value={lazyRadius}
            onChange={(e) => ( updateLazyRadius(parseInt(e.target.value, 10)) )}
          />
        </div>
      </div>
      <CanvasDraw
        ref={(canvasDraw: any) => (saveableCanvas = canvasDraw)}
        brushColor={color}
        brushRadius={brushRadius}
        lazyRadius={lazyRadius}
        canvasWidth={width}
        canvasHeight={height}
        onChange={() => console.log('onChange')}
      />
    </div>
  );
}

export default PrintPanel;