import React from "react";
import InvoiceSentCard from "./InvoiceSentCard";

const InvoiceSent = () => {
  return (
    <div className="md:w-4/12">
      <h1 className="text-[#333B69] py-2 font-semibold">Invoices Sent</h1>
      <div className=" max-w-sm  bg-white border border-gray-200 rounded-[15px] shadow p-4">
        <div className="flow-root">
          <ul role="list" className=" ">
            <li className="py-1">
              <InvoiceSentCard
                name="Apple Store"
                time="5h"
                amount={450}
                imageUrl="/assets/images/apple.png"
              />
            </li>
            <li className="py-1">
              <InvoiceSentCard
                name="Michael"
                time="2 days"
                amount={160}
                imageUrl="/assets/images/person.png"
              />
            </li>
            <li className="py-1">
              <InvoiceSentCard
                name="Playstation"
                time="5 days"
                amount={1085}
                imageUrl="/assets/images/playstation.png"
              />
            </li>
            <li className="py-1">
              <InvoiceSentCard
                name="William"
                time="10 days"
                amount={90}
                imageUrl="/assets/images/person2.png"
              />
            </li>
          </ul>
        </div>
      </div>
    </div>
  );
};

export default InvoiceSent;
