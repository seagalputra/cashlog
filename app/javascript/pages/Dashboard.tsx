import React from "react";
import Main from "../layouts/Main";

import { Header, List } from "../components";

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

      <List title="Recent Transactions">
        <List.Item>
          <div className="flex gap-5 items-center">
            <List.Icon width="50px" height="50px">
              T
            </List.Icon>

            <div className="flex flex-col gap-1">
              <p className="font-sans text-lg font-bold">Freelance</p>
              <p className="font-sans text-gray-400 text-md">Aug 10, 2020</p>
            </div>
          </div>

          <div className="flex flex-col items-end gap-1">
            <p className="font-sans text-lg font-bold text-green-500">
              Rp 50.000
            </p>
            <p className="font-sans text-gray-400 text-md">Income</p>
          </div>
        </List.Item>
      </List>
    </Main>
  );
};
