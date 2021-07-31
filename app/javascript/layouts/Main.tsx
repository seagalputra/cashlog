import React from "react";
import { Navbar } from "../components";

export default ({ children }) => {
  return (
    <>
      <Navbar className="bg-gray-800">
        <div className="container flex justify-between items-center mx-auto md:max-w-lg md:px-3">
          <Navbar.Title>
            <p className="font-bold text-white">Cashlog</p>
          </Navbar.Title>
          <Navbar.Item>
            <button
              type="button"
              className="max-w-xs bg-gray-800 rounded-full flex items-center text-sm focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-gray-800 focus:ring-white"
            >
              <span className="sr-only">Open user menu</span>
              <img
                className="h-8 w-8 rounded-full"
                src="https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80"
                alt=""
              />
            </button>
          </Navbar.Item>
        </div>
      </Navbar>

      <main className="min-h-screen pt-16 bg-gray-100">
        <div className="px-3 py-4 container mx-auto md:max-w-lg">
          {children}
        </div>
      </main>
    </>
  );
};
