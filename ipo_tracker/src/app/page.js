// // import Image from "next/image";

// // export default function Home() {
// //   return (
// //     <div className="grid grid-rows-[20px_1fr_20px] items-center justify-items-center min-h-screen p-8 pb-20 gap-16 sm:p-20 font-[family-name:var(--font-geist-sans)]">
// //       <main className="flex flex-col gap-8 row-start-2 items-center sm:items-start">
// //         <Image
// //           className="dark:invert"
// //           src="/next.svg"
// //           alt="Next.js logo"
// //           width={180}
// //           height={38}
// //           priority
// //         />
// //         <ol className="list-inside list-decimal text-sm text-center sm:text-left font-[family-name:var(--font-geist-mono)]">
// //           <li className="mb-2">
// //             Get started by editing{" "}
// //             <code className="bg-black/[.05] dark:bg-white/[.06] px-1 py-0.5 rounded font-semibold">
// //               src/app/page.js
// //             </code>
// //             .
// //           </li>
// //           <li>Save and see your changes instantly.</li>
// //         </ol>

// //         <div className="flex gap-4 items-center flex-col sm:flex-row">
// //           <a
// //             className="rounded-full border border-solid border-transparent transition-colors flex items-center justify-center bg-foreground text-background gap-2 hover:bg-[#383838] dark:hover:bg-[#ccc] text-sm sm:text-base h-10 sm:h-12 px-4 sm:px-5"
// //             href="https://vercel.com/new?utm_source=create-next-app&utm_medium=appdir-template-tw&utm_campaign=create-next-app"
// //             target="_blank"
// //             rel="noopener noreferrer"
// //           >
// //             <Image
// //               className="dark:invert"
// //               src="/vercel.svg"
// //               alt="Vercel logomark"
// //               width={20}
// //               height={20}
// //             />
// //             Deploy now
// //           </a>
// //           <a
// //             className="rounded-full border border-solid border-black/[.08] dark:border-white/[.145] transition-colors flex items-center justify-center hover:bg-[#f2f2f2] dark:hover:bg-[#1a1a1a] hover:border-transparent text-sm sm:text-base h-10 sm:h-12 px-4 sm:px-5 sm:min-w-44"
// //             href="https://nextjs.org/docs?utm_source=create-next-app&utm_medium=appdir-template-tw&utm_campaign=create-next-app"
// //             target="_blank"
// //             rel="noopener noreferrer"
// //           >
// //             Read our docs
// //           </a>
// //         </div>
// //       </main>
// //       <footer className="row-start-3 flex gap-6 flex-wrap items-center justify-center">
// //         <a
// //           className="flex items-center gap-2 hover:underline hover:underline-offset-4"
// //           href="https://nextjs.org/learn?utm_source=create-next-app&utm_medium=appdir-template-tw&utm_campaign=create-next-app"
// //           target="_blank"
// //           rel="noopener noreferrer"
// //         >
// //           <Image
// //             aria-hidden
// //             src="/file.svg"
// //             alt="File icon"
// //             width={16}
// //             height={16}
// //           />
// //           Learn
// //         </a>
// //         <a
// //           className="flex items-center gap-2 hover:underline hover:underline-offset-4"
// //           href="https://vercel.com/templates?framework=next.js&utm_source=create-next-app&utm_medium=appdir-template-tw&utm_campaign=create-next-app"
// //           target="_blank"
// //           rel="noopener noreferrer"
// //         >
// //           <Image
// //             aria-hidden
// //             src="/window.svg"
// //             alt="Window icon"
// //             width={16}
// //             height={16}
// //           />
// //           Examples
// //         </a>
// //         <a
// //           className="flex items-center gap-2 hover:underline hover:underline-offset-4"
// //           href="https://nextjs.org?utm_source=create-next-app&utm_medium=appdir-template-tw&utm_campaign=create-next-app"
// //           target="_blank"
// //           rel="noopener noreferrer"
// //         >
// //           <Image
// //             aria-hidden
// //             src="/globe.svg"
// //             alt="Globe icon"
// //             width={16}
// //             height={16}
// //           />
// //           Go to nextjs.org →
// //         </a>
// //       </footer>
// //     </div>
// //   );
// // }
// "use client"; // Ensure this is a client component

// import { useEffect, useState } from "react";

// // export default function Home() {
// //   const [data, setData] = useState(null);
// //   const [error, setError] = useState(null);

// //   useEffect(() => {
// //     const fetchData = async () => {
// //       try {
// //         const response = await fetch("http://localhost:8080/data", { mode: "no-cors" }); // Replace with your API endpoint
// //         if (!response.ok) {
// //           throw new Error(`HTTP error! status: ${response.status}`);
// //         }
// //         const result = await response.json();
// //         console.log(result)
// //         setData(result); // Save the response
// //       } catch (err) {
// //         setError(err.message); // Handle any errors
// //       }
// //     };

// //     fetchData(); // Call the function
// //   }, []); // Run only once when the component mounts

// //   return (
// //     <div>
// //       <h1>Response</h1>
// //       {error && <p>Error: {error}</p>}
// //       {data ? <pre>{JSON.stringify(data, null, 2)}</pre> : <p>Loading...</p>}
// //     </div>
// //   );
// // }


// const MyPage = () => {
//   const [data, setData] = useState([]);

//   useEffect(() => {
//     // Function to fetch data
//     const fetchData = async () => {
//       const response = await fetch('http://0.0.0.0:8080/data', { mode: "no-cors" }, { cache: 'no-store' }); // Replace with your API endpoint
//       const result = await response.json();
//       setData(result);
//     };

//     // Call the function
//     fetchData();
//   }, []); // Empty array ensures this effect runs once on mount


//   return (
//     <div className="p-20">
//       <h1 className="text-2xl font-bold mb-4">Upcoming IPO Table</h1>
//       <div className="overflow-x-auto">
//         {/* Render your data here */}
//         {/* {data && <pre>{JSON.stringify(data, null, 2)}</pre>}\ */}
//         {data && (
//         <table className="table-auto border-collapse border border-gray-300 w-full text-left">
//           <thead>
//             <tr>
//               {data.headers.map((header, index) => (
//                 <th className="border border-gray-300 px-4 py-2" key={index}>{header}</th>
//               ))}
//             </tr>
//           </thead>
//           <tbody>
//             {data.rows.map((row, rowIndex) => (
//               <tr key={rowIndex}>
//                 {row.map((cell, cellIndex) => (
//                   <td className="border border-gray-300 px-4 py-2" key={cellIndex}>{cell}</td>
//                 ))}
//               </tr>
//             ))}
//           </tbody>
//         </table>
//       )}
//       </div>
//     </div>
//   );
// };

// export default MyPage;

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