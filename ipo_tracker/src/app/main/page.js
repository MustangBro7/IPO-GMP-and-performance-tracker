"use client";

import { useEffect, useState } from "react";

const extractPercentage = (content) => {
  const start = content.lastIndexOf("(");
  const end = content.lastIndexOf("%");
  if (start !== -1 && end !== -1 && end > start) {
    let percentageStr = content.slice(start + 1, end);
    percentageStr = percentageStr.replace(/,/g, "");
    const percentage = parseFloat(percentageStr);
    return isNaN(percentage) ? 0 : percentage;
  }
  return 0;
};

const MyPage = () => {
  const [data, setData] = useState(null); 
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch("http://localhost:8080/main");
        if (!response.ok) {
          throw new Error(`HTTP error! Status: ${response.status}`);
        }
        const result = await response.json();
        setData(result);
      } catch (err) {
        console.error("Error fetching data:", err);
        setError(err.message);
      }
    };

    fetchData();
  }, []);

  return (
    <div className="p-20">
      <h1 className="text-2xl font-bold mb-4">Mainboard IPO Performace Table</h1>
      <div className="overflow-x-auto">
        {error ? (
          <p className="text-red-500">Error: {error}</p>
        ) : data ? (
          <table className="table-auto border-collapse border border-gray-300 w-full text-left">
            <thead>
              <tr>
                {data.headers.map((header, index) => (
                  <th key={index} className="border border-gray-300 px-4 py-2 font-bold">
                    {header}
                  </th>
                ))}
              </tr>
            </thead>
            <tbody>
              {data.rows.map((row, rowIndex) => (
                <tr key={rowIndex}>
                  {row.map((cell, cellIndex) => {
                    if (typeof cell === "string") {
                      if (cell.includes("SME")) {
                        return (
                          <td
                            key={cellIndex}
                            className="border border-gray-300 px-4 py-2 bg-yellow-200 dark:bg-purple-700"
                          >
                            {cell}
                          </td>
                        );
                      }

                      const percentage = extractPercentage(cell);
                      if (percentage > 0) {
                        return (
                          <td
                            key={cellIndex}
                            className="border border-gray-300 px-4 py-2 bg-green-500"
                          >
                            {cell}
                          </td>
                        );
                      } else if (percentage < 0) {
                        return (
                          <td
                            key={cellIndex}
                            className="border border-gray-300 px-4 py-2 bg-red-500"
                          >
                            {cell}
                          </td>
                        );
                      }
                    }
                    return (
                      <td
                        key={cellIndex}
                        className="border border-gray-300 px-4 py-2"
                      >
                        {cell}
                      </td>
                    );
                  })}
                </tr>
              ))}
            </tbody>
          </table>
        ) : (
          <p className="text-gray-500">Loading...</p>
        )}
      </div>
    </div>
  );
};

export default MyPage;
