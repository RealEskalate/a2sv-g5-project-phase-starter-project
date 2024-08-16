import React from "react";
import LoanCards from "./LoanCards";
import LoanTable from "./LoanTable";

function page() {
	return (
		<div className="px-10 py-8">
			<LoanCards />
			<LoanTable />
		</div>
	);
}

export default page;
