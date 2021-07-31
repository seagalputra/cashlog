import React from "react";
import { Link } from "@inertiajs/inertia-react";
import Main from "../layouts/Main";

export default ({ recent_transactions, summary }) => {
  const { needs, wants, invest } = summary;

  return (
    <Main>
      <div className="w-full p-4 bg-white border mb-2">
        <div className="flex flex-col gap-2">
          <p className="font-light text-gray-400">Your current balance</p>
          <p className="text-2xl font-bold">{needs}</p>
        </div>
      </div>

      <div className="flex justify-between gap-2">
        <div className="w-full p-4 bg-white border">
          <div className="flex flex-col gap-2">
            <p className="text-lg font-bold">{needs}</p>
            <p className="font-light text-gray-400">Needs</p>
          </div>
        </div>
        <div className="w-full p-4 bg-white border">
          <div className="flex flex-col gap-2">
            <p className="text-lg font-bold">{wants}</p>
            <p className="font-light text-gray-400">Wants</p>
          </div>
        </div>
        <div className="w-full p-4 bg-white border">
          <div className="flex flex-col gap-2">
            <p className="text-lg font-bold">{invest}</p>
            <p className="font-light text-gray-400">Invest</p>
          </div>
        </div>
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
