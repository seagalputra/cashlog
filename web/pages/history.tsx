import { NextPage } from "next";
import Head from "next/head";
import Link from "next/link";

// TODO: Place ListItem to components folder
import { ListItem } from "./index";

const TransactionHistory: NextPage = () => {
	return (
		<>
			<Head>
				<title>History | Cashlog</title>
				<link rel="icon" href="/favicon.ico" />
			</Head>

			<main>
				<div className="flex justify-start items-center pt-16 max-w-xl w-full mx-auto">
					<Link href="/">
						<a
							className="flex justify-center items-center rounded-full border-2 mr-6"
							style={{
								borderColor: "#18191F",
								height: "64px",
								width: "64px",
								boxShadow: "0px 2px 0px #18191F",
							}}
						>
							<svg
								xmlns="http://www.w3.org/2000/svg"
								width="24"
								height="24"
								fill="none"
								stroke="currentColor"
								strokeWidth="2"
								strokeLinecap="round"
								strokeLinejoin="round"
								className="feather feather-chevron-left"
							>
								<path d="M15 18l-6-6 6-6" />
							</svg>
						</a>
					</Link>
					<p
						className="font-extrabold text-4xl "
						style={{
							fontFamily: "Montserrat, sans-serif",
						}}
					>
						Transaction
					</p>
				</div>
				<div className="flex justify-center mt-8">
					<input
						className="border-2 py-4 pl-6 rounded-2xl w-full max-w-xl"
						style={{ borderColor: "#18191F", boxShadow: "0px 2px 0px #18191F" }}
						type="text"
						placeholder="Search transaction"
					/>
				</div>
				{["1", "2", "3", "4"].map((item) => (
					<ListItem key={item} />
				))}
			</main>
		</>
	);
};

export default TransactionHistory;
