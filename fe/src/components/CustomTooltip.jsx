import React from 'react'
import { formatTime } from '../utils/utils';

const CustomTooltip = ({ active, payload }) => {
    if (active && payload && payload.length) {
      const { status_code, route, response_time, method, created_at } = payload[0].payload;
      return (
        <div style={{
          backgroundColor: "#fff", 
          border: "1px solid #ccc", 
          padding: "10px", 
          borderRadius: "5px",
          boxShadow: "0px 4px 6px rgba(0, 0, 0, 0.1)"
        }}
        >
          <p><strong>Route:</strong>{route}</p>
          <p><strong>Method:</strong> {method}</p>
          <p><strong>Name:</strong> {status_code}</p>
          <p><strong>Response Time:</strong> {response_time} ms</p>
          <p><strong>Created At:</strong> {formatTime(created_at)}</p>
        </div>
      );
    }

    return null;
};

export default CustomTooltip