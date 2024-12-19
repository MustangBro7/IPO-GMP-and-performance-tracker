"use client";

import { useEffect, useState } from "react";

const MyPage = () => {
  const [data, setData] = useState(null); // Start with null since the data structure is an object
  const [error, setError] = useState(null); // Track any errors

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch("http://localhost:8080/data");
        if (!response.ok) {
          throw new Error(`HTTP error! Status: ${response.status}`);
        }
        const result = await response.json(); // Expecting { headers: [], rows: [] }
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
      <h1 className="text-2xl font-bold mb-4">IPO Table</h1>
      <div className="overflow-x-auto">
        {error ? (
          <p className="text-red-500">Error: {error}</p>
        ) : data ? (
          <table className="table-auto border-collapse border border-gray-300 w-full text-left">
            <thead>
              <tr className="">
                {data.headers.map((header, index) => (
                  <th
                    key={index}
                    className="border border-gray-300 px-4 py-2 font-bold"
                  >
                    {header}
                  </th>
                ))}
              </tr>
            </thead>
            <tbody>
              {data.rows.length > 0 ? (
                data.rows.map((row, index) => (
                  <tr
                    key={index}
                    className={index % 2 === 0 ? "" : ""}
                  >
                    {row.map((cell, cellIndex) => (
                      <td
                        key={cellIndex}
                        className="border border-gray-300 px-4 py-2"
                      >
                        {cell}
                      </td>
                    ))}
                  </tr>
                ))
              ) : (
                <tr>
                  <td
                    colSpan={data.headers.length}
                    className="text-center border border-gray-300 px-4 py-2"
                  >
                    No data available
                  </td>
                </tr>
              )}
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
