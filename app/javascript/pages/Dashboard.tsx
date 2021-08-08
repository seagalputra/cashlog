import React from "react";
import Main from "../layouts/Main";

import { Summary, List, Menu, Icon } from "../components";

export default ({ recent_transactions, summary }) => {
  return (
    <Main>
      <Summary className="mb-2">
        <Summary.Item>
          <Summary.Title>Your current balance</Summary.Title>
          <p className="text-2xl font-bold">0</p>
        </Summary.Item>
      </Summary>

      <div className="flex justify-between gap-2">
        {Object.entries(summary).map((obj, index) => {
          const [key, value]: any = obj;
          const [first, ...rest] = key;
          return (
            <Summary key={`header-${index}`}>
              <Summary.Item>
                <p className="text-lg font-bold">{value}</p>
                <Summary.Title>
                  {first.toUpperCase() + rest.join("")}
                </Summary.Title>
              </Summary.Item>
            </Summary>
          );
        })}
      </div>

      <Menu>
        <Menu.Item>
          <Icon.Wallet
            className="flex items-center justify-center h-10 w-10 border rounded-full"
            aria-hidden="true"
          />
          <Menu.Link to="#">Transaction</Menu.Link>
        </Menu.Item>

        <Menu.Item>
          <Icon.File
            className="flex items-center justify-center h-10 w-10 border rounded-full"
            aria-hidden="true"
          />
          <Menu.Link to="#">History</Menu.Link>
        </Menu.Item>
      </Menu>

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
