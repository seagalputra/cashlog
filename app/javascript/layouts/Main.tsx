import React from "react";
import Navbar from "../components/Navbar";

export default ({ children }) => {
  return (
    <>
      <Navbar />
      <main className="min-h-screen pt-16 bg-gray-100">
        <div className="px-6 py-4 container mx-auto md:max-w-lg">
          {children}
        </div>
      </main>
    </>
  );
};
