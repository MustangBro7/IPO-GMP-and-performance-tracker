"use client";

import { useEffect, useState } from "react";
import Link from "next/link";

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
  const [loading, setLoading] = useState(false);

  const fetchData = async () => {
    const apiUrl = process.env.NEXT_PUBLIC_API_URL;
    setLoading(true);
    try {
      const response = await fetch(`${apiUrl}/data`, {
        method: "GET",
        headers: {
          "ngrok-skip-browser-warning": "true",
        },
      });
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }
      const result = await response.json();
      setData(result);
    } catch (err) {
      console.error("Error fetching data:", err);
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  const loadingFunc = () => {
    setLoading(true);
  };

  useEffect(() => {
    fetchData();
  }, []);

  return (
    <div className="p-4 sm:p-8">
      {/* Navigation Buttons */}
      <div>
        <Link href="/">
          <button 
          // onClick={loadingFunc}
          className="px-4 py-2 bg-black-500  rounded hover:bg-gray-600 outline rounded-b-none  outline-1">
          <h1 className="text-2xl font-bold mb-4">Upcoming IPO Table</h1>
          </button>
        </Link>
        <Link href="/main">
          <button 
          onClick={loadingFunc}
          className="px-4 py-2 bg-black-500  rounded hover:bg-gray-600">
          <h1 className="text-2xl font-bold mb-4">Main</h1>
          </button>
        </Link>
        <Link href="/sme">
          <button 
          onClick={loadingFunc}
          className="px-4 py-2 bg-black-500  rounded hover:bg-gray-600">
          <h1 className="text-2xl font-bold mb-4 ">SME</h1>
          </button>
        </Link>
        </div>

      {/* Loading Overlay */}
      {loading && (
        <div className="absolute inset-0 flex items-center justify-center bg-black bg-opacity-50 z-50">
          <div className="w-16 h-16 border-4 border-t-blue-500 border-gray-300 rounded-full animate-spin"></div>
        </div>
      )}

      {/* Main Content */}
      <div className="overflow-x-auto">
        {error ? (
          <p className="text-red-500">Error: {error}</p>
        ) : data ? (
          <table className="table-auto border-collapse border border-gray-300 w-full text-left sm:table-fixed">
            <thead>
              <tr>
                {data.headers.map((header, index) => (
                  <th
                    key={index}
                    className="border border-gray-300 px-2 sm:px-4 py-1 sm:py-2 font-bold text-xs sm:text-sm"
                  >
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
                            className="border border-gray-300 px-2 sm:px-4 py-1 sm:py-2 bg-yellow-200 dark:bg-purple-700 text-xs sm:text-sm"
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
                            className="border border-gray-300 px-2 sm:px-4 py-1 sm:py-2 bg-green-500 text-xs sm:text-sm"
                          >
                            {cell}
                          </td>
                        );
                      } else if (percentage < 0) {
                        return (
                          <td
                            key={cellIndex}
                            className="border border-gray-300 px-2 sm:px-4 py-1 sm:py-2 bg-red-500 text-xs sm:text-sm"
                          >
                            {cell}
                          </td>
                        );
                      }
                    }
                    return (
                      <td
                        key={cellIndex}
                        className="border border-gray-300 px-2 sm:px-4 py-1 sm:py-2 text-xs sm:text-sm"
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
          <p className="text-gray-500 text-sm sm:text-base">Loading...</p>
        )}
      </div>
    </div>
  );
};

export default MyPage;
