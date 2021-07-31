import React from "react";

export default () => {
  return (
    <nav className="fixed inset-x-0 top-0 z-50 flex items-center justify-between w-full p-3 bg-gray-800 ">
      <div className="flex items-center space-x-6">
        <p className="font-bold text-white">Cashlog</p>
      </div>

      <div className="flex items-center gap-2">
        <div className="ml-3 relative">
          <div>
            <button
              type="button"
              className="max-w-xs bg-gray-800 rounded-full flex items-center text-sm focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-gray-800 focus:ring-white"
              id="user-menu-button"
              aria-expanded="false"
              aria-haspopup="true"
            >
              <span className="sr-only">Open user menu</span>
              <img
                className="h-8 w-8 rounded-full"
                src="https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80"
                alt=""
              />
            </button>
          </div>

          <div
            id="user-menu"
            className="origin-top-right absolute right-0 mt-2 w-48 rounded-md shadow-lg py-1 bg-white ring-1 ring-black ring-opacity-5 focus:outline-none invisible"
            role="menu"
            aria-orientation="vertical"
            aria-labelledby="user-menu-button"
            tabIndex={-1}
          >
            <a
              href="#"
              className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
              role="menuitem"
              tabIndex={-1}
              id="user-menu-item-0"
            >
              Your Profile
            </a>

            <a
              href="#"
              className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
              role="menuitem"
              tabIndex={-1}
              id="user-menu-item-1"
            >
              Settings
            </a>

            <a
              href="#"
              className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
              role="menuitem"
              tabIndex={-1}
              id="user-menu-item-2"
            >
              Sign out
            </a>
          </div>
        </div>
      </div>
    </nav>
  );
};
