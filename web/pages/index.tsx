import Head from "next/head";

export default function Home() {
	return (
		<>
			<Head>
				<title>Home | Cashlog</title>
				<link rel="icon" href="/favicon.ico" />
			</Head>

			<main>
				<div
					className="flex justify-center h-80 w-full pt-4 relative"
					style={{
						backgroundColor: "#00C6AE",
					}}
				>
					<img src="/person.svg" alt="Two person" />
					<div
						className="bg-white border-2 p-6 absolute w-full max-w-xl rounded-2xl"
						style={{
							bottom: "-130px",
							borderColor: "#18191F",
						}}
					>
						<p
							className="text-lg mb-2"
							style={{
								fontFamily: "Montserrat, sans-serif",
								fontWeight: 500,
							}}
						>
							Current Balance
						</p>
						<p
							className="text-4xl"
							style={{
								fontFamily: "Montserrat, sans-serif",
								fontWeight: "bold",
								color: "#18191F",
							}}
						>
							Rp 1.500.000
						</p>
						<div className="flex flex-row mt-9">
							<button
								className="flex flex-row justify-center content-center p-4 border-2 rounded-2xl"
								style={{
									backgroundColor: "#FFBD12",
									borderColor: "#18191F",
									boxSizing: "border-box",
									boxShadow: "0px 4px 0px #18191F",
								}}
							>
								<svg
									width="24"
									height="24"
									fill="none"
									xmlns="http://www.w3.org/2000/svg"
								>
									<path
										d="M9.12 2.88a2.88 2.88 0 115.76 0v18.24a2.88 2.88 0 11-5.76 0V2.88z"
										fill="#18191F"
									/>
									<path
										d="M2.88 14.88a2.88 2.88 0 110-5.76h18.24a2.88 2.88 0 110 5.76H2.88z"
										fill="#18191F"
									/>
								</svg>
								<p
									className="ml-2"
									style={{
										fontFamily: "Monserrat, sans-serif",
										fontWeight: 800,
									}}
								>
									Add Transaction
								</p>
							</button>
							<button
								className="p-4 border-2 rounded-2xl ml-4"
								style={{
									borderColor: "#18191F",
									boxSizing: "border-box",
									boxShadow: "0px 4px 0px #18191F",
								}}
							>
								<svg
									xmlns="http://www.w3.org/2000/svg"
									width="24"
									height="24"
									fill="none"
									stroke="currentColor"
									stroke-width="2"
									stroke-linecap="round"
									stroke-linejoin="round"
									className="feather feather-archive"
								>
									<path d="M21 8v13H3V8M1 3h22v5H1zM10 12h4" />
								</svg>
							</button>
						</div>
					</div>
				</div>
			</main>
		</>
	);
}
