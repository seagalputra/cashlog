import React from "react";
import { Link } from "@inertiajs/inertia-react";
import Main from "../layouts/Main";

import { Header } from "../components/Header";

export default ({ recent_transactions, summary }) => {
  return (
    <Main>
      <Header className="mb-2">
        <Header.Item>
          <p className="font-light text-gray-400">Your current balance</p>
          <p className="text-2xl font-bold">0</p>
        </Header.Item>
      </Header>

      <div className="flex justify-between gap-2">
        {Object.entries(summary).map((obj, index) => {
          const [key, value]: any = obj;
          const [first, ...rest] = key;
          return (
            <Header key={`header-${index}`}>
              <Header.Item>
                <p className="text-lg font-bold">{value}</p>
                <p className="font-light text-gray-400">
                  {first.toUpperCase() + rest.join("")}
                </p>
              </Header.Item>
            </Header>
          );
        })}
      </div>

      <div className="flex flex-row items-center mt-6 mb-6 gap-4">
        <p className="text-xl font-semibold text-gray-400 flex-0">
          Recent Transactions
        </p>
        <div className="flex-grow border-b"></div>
        <Link href="#" className="text-gray-400 hover:text-gray-500">
          View more
        </Link>
      </div>
    </Main>
  );
};
