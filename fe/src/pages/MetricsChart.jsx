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
    }).catch((err) =>
      console.log("Failed to get data:", err)
    );
    
    const socket = new WebSocket(METRICS_WEBSOCKET_URL);

    socket.onmessage = (event) => {
      try {
        const newMetric = JSON.parse(event.data);
        newMetric.created_at = new Date(newMetric.created_at).getTime();
        setMetrics((prev) => [...prev, newMetric].slice(-600));
      } catch (err) {
        console.error('Parsing data error:', err);
      }
    };

    socket.onerror = (error) => {
      console.log('Websocket connection error:', error);
    };

    socket.onclose = (event) => {
      if (!event.wasClean) {
        setTimeout(() => {
          window.location.reload();
        }, 3000);
      }
    };

    return () => {
        socket.close();
    }
  }, []);

  return (
    <div class="chart">
      {
        metrics.length === 0 ? 
        <>
          No data to display, something went wrong
        </> :
        <>
          <LineChart width={1200} height={300} data={metrics}>
            <CartesianGrid strokeDasharray="4 4" />
            <XAxis 
              dataKey="created_at"
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
        </>
      }
    </div>
  );
};

export default MetricsChart;
