import React, { useEffect, useState } from "react";
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip } from 'recharts';
import { formatTime } from "../utils/utils";
import CustomTooltip from "../components/CustomTooltip";
import { METRICS_API_URL, METRICS_WEBSOCKET_URL } from "../utils/constants";

const MetricsChart = () => {
  const [metrics, setMetrics] = useState([]);

  useEffect(() => {
    fetch(METRICS_API_URL)
      .then((res) => res.json())
      .then((data) => {
         const formattedData = data.map((item) => ({
          ...item,
          created_at: new Date(item.created_at).getTime(),
        }));
        setMetrics(formattedData);
    });
    
    const socket = new WebSocket(METRICS_WEBSOCKET_URL);

    socket.onmessage = (event) => {
      const newMetric = JSON.parse(event.data);
      newMetric.created_at = new Date(newMetric.created_at).getTime();
      setMetrics((prev) => [...prev, newMetric].slice(-600));
    };

    return () => {
        socket.close();
    }
  }, []);

  // const now = Date.now();
  // const tenMinutesAgo = now - 10 * 60 * 1000;

  return (
    <div style={{
      width: '100%',
      height: '64%',
      margin: '50px'
    }}>
      <LineChart width={1000} height={300} data={metrics}>
        <CartesianGrid strokeDasharray="4 4" />
        <XAxis 
          dataKey="created_at"
          // type="number"
          // domain={[tenMinutesAgo, now]}
          tickFormatter={formatTime}     
        />
        <YAxis />
        <Tooltip 
          content={<CustomTooltip />}
        />
        <Line 
          type="monotone"
          dataKey="response_time"
          stroke="#82ca9d"
          dot={false}
          animationDuration={10}
        />
      </LineChart>
    </div>
  );
};

export default MetricsChart;
