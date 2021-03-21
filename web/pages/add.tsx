import { NextPage } from "next";
import Head from "next/head";
import Link from "next/link";

const AddTransaction: NextPage = () => {
	return (
		<>
			<Head>
				<title>Add Transaction | Cashlog</title>
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
						className="font-extrabold text-4xl"
						style={{
							fontFamily: "Montserrat, sans-serif",
						}}
					>
						Add Transaction
					</p>
				</div>
				<form className="flex flex-col mt-20 mx-auto w-full max-w-xl">
					<input
						className="border-b-2 py-4 pl-6 text-2xl font-bold"
						style={{
							borderColor: "#18191F",
							fontFamily: "Montserrat, sans-serif",
						}}
						type="text"
						placeholder="Enter amount"
					/>
					<p
						className="mt-12 font-bold text-xl"
						style={{ fontFamily: "Montserrat, sans-serif " }}
					>
						Select category
					</p>
					<div className="flex justify-center content-center mt-6">
						<select
							className="appearance-none border-t-2 border-b-2 border-l-2 p-4 rounded-tl-2xl rounded-bl-2xl w-full"
							style={{
								borderColor: "#18191F",
							}}
						>
							<option>Needs</option>
							<option>Wants</option>
							<option>Invest</option>
						</select>
						<button
							className="border-t-2 border-r-2 border-b-2 p-4 rounded-tr-2xl rounded-br-2xl"
							style={{
								borderColor: "#18191F",
								backgroundColor: "#FFBD12",
							}}
							onClick={(event) => event.preventDefault()}
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
								className="feather feather-chevron-down"
							>
								<path d="M6 9l6 6 6-6" />
							</svg>
						</button>
					</div>
					<p
						className="mt-12 font-bold text-xl"
						style={{ fontFamily: "Montserrat, sans-serif " }}
					>
						Select transaction status
					</p>
					<div className="flex mt-6 gap-4">
						<div
							className="border-2 py-2 px-4 rounded-full font-bold cursor-pointer"
							style={{
								borderColor: "#18191F",
								fontFamily: "Montserrat, sans-serif",
							}}
						>
							<p>Income</p>
						</div>
						<div
							className="border-2 py-2 px-4 rounded-full font-bold cursor-pointer"
							style={{
								color: "#18191F",
								borderColor: "#18191F",
								backgroundColor: "#FFBD12",
								fontFamily: "Montserrat, sans-serif",
							}}
						>
							<p>Outcome</p>
						</div>
						<div
							className="border-2 py-2 px-4 rounded-full font-bold cursor-pointer"
							style={{
								borderColor: "#18191F",
								fontFamily: "Montserrat, sans-serif",
							}}
						>
							<p>Waiting</p>
						</div>
					</div>
					<p
						className="mt-12 font-bold text-xl"
						style={{ fontFamily: "Montserrat, sans-serif " }}
					>
						Description
					</p>
					<textarea
						className="border-2 py-4 pl-6 rounded-2xl mt-6 h-64"
						style={{
							borderColor: "#18191F",
							fontFamily: "Montserrat, sans-serif",
						}}
						placeholder="Write a description"
					/>
					<button
						className="my-12 p-4 w-full border-2 text-white rounded-2xl font-extrabold text-xl"
						onClick={(event) => event.preventDefault()}
						style={{
							fontFamily: "Montserrat, sans-serif",
							backgroundColor: "#18191F",
						}}
					>
						Submit
					</button>
				</form>
			</main>
		</>
	);
};

export default AddTransaction;
