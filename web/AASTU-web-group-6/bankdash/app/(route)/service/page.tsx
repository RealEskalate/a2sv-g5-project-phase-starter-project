"use client";
import React, { useEffect, useState } from "react";
import DescriptionCard from "@/app/components/Card/DescriptionCard";
import ServicesCard from "@/app/components/Card/ServicesCard";
import axios from "axios";
import { useSession } from "next-auth/react";
import ModalService from "@/app/components/Card/ModalService";
import ShimmerCard from "@/app/components/Shimmer/SimmerCard";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faAdd } from "@fortawesome/free-solid-svg-icons";

interface BankService {
  id: string;
  name: string;
  details: string;
  numberOfUsers: number;
  status: string;
  type: string;
  icon: string;
  colors: string;
}

const Services = () => {
  const { data: session } = useSession();
  const [services, setServices] = useState<BankService[]>([]);
  const [pageNumber, setPageNumber] = useState(1);
  const [isModalOpen, setIsModalOpen] = useState(false);

  const handleModalToggle = () => {
    setIsModalOpen(!isModalOpen);
  };

  const accessToken = session?.accessToken as string;
  const li = [1, 2, 3, 4];

  async function fetchData(accessToken: string) {
    try {
      const response = await axios.get(
        `https://bank-dashboard-rsf1.onrender.com/bank-services?page=0&size=50`,
        {
          headers: {
            Authorization: `Bearer ${accessToken}`,
          },
        }
      );
      setServices(response.data.data.content);
      console.log(services);
    } catch (error) {
      console.error("There was a problem with the axios request:", error);
    }
  }

  useEffect(() => {
    fetchData(accessToken);
  }, [accessToken]);

  return (
    <div className="w-[100%] xxs:pt-10 xs:pt-12 md:pt-12 lg:pt-0 px-6">
      <div className="">
        <div className=" lg:mr-0 flex gap-4 overflow-x-auto lg:overflow-x-visible [&::-webkit-scrollbar]:hidden [-ms-overflow-style:none] [scrollbar-width:none] lg:pt-10 w-full">
          {/* <div className="w-[100%] lg:w-[350px] flex-shrink-0"> */}
          <ServicesCard
            img="/assets/lifeInsurance.svg"
            title="Life Insurance"
            desc="Unlimited Protection"
          />
          {/* </div> */}
          {/* <div className="w-[100%] lg:w-[350px] flex-shrink-0"> */}
          <ServicesCard
            img="/assets/shoppingBag.svg"
            title="Shopping"
            desc="Buy. Think. Grow"
          />
          {/* </div> */}
          {/* <div className="w-[100%] lg:w-[350px] flex-shrink-0"> */}
          <ServicesCard
            img="/assets/safety.svg"
            title="Safety"
            desc="We are your allies"
          />
        </div>
        {/* </div> */}

        <div>
          <div className="flex justify-between">
            <p className="font-semibold xxs:text-lg md:text-[22px] text-[#343C6A] dark:text-gray-300 pt-5 pb-5 lg:p-10 ">
              Bank Services List
            </p>
            <div
              className={`flex items-center text-base text-[#718EBF] dark:text-gray-400 rounded-[50px] py-1 pl-6 grow justify-end ${
                isModalOpen ? "blur-sm" : ""
              }`}
            >
              <button
                onClick={handleModalToggle}
                className="flex gap-2 items-center text-white p-3 px-6 rounded-xl shadow-md bg-[#1814F3] hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500"
                type="button"
              >
                <FontAwesomeIcon icon={faAdd} className="font-semibold" />
                Add
              </button>
            </div>
            {isModalOpen && (
              <div
                className="fixed inset-0 z-50 flex justify-center items-center bg-black dark:bg-white dark:bg-opacity-[4%] dark:border-white dark:border-[1px] bg-opacity-50 backdrop-blur-sm"
                onClick={handleModalToggle}
              >
                <div
                  className="relative bg-white dark:bg-[#232328] p-6 rounded-lg shadow-lg max-w-lg w-full"
                  onClick={(e) => e.stopPropagation()} // Prevent modal from closing when clicking inside it
                >
                  <ModalService
                    isOpen={isModalOpen}
                    onClose={handleModalToggle}
                  />
                </div>
              </div>
            )}
          </div>
          <div className="w-full flex flex-col grow items-start">
            {services.length > 0 ? (
              services.map((service, index) => (
                <>
                  {console.log(service.icon)}
                  <DescriptionCard
                    key={service.id}
                    img={service.icon}
                    title={service.name}
                    desc={service.details}
                    colOne="Number of Users"
                    descOne={service.numberOfUsers}
                    colTwo="Status"
                    descTwo={service.status}
                    colThree="Type"
                    descThree={service.type}
                    btn="View Details"
                    color={
                      // Split the URL at the '?' to remove query parameters, then check for the file name
                      service.icon.includes("services2.svg")
                        ? "bg-orange-100"
                        : service.icon.includes("services4.svg")
                        ? "bg-blue-100"
                        : service.icon.includes("services5.svg")
                        ? "bg-green-100"
                        : "bg-pink-100"
                    }
                  />
                </>
              ))
            ) : (
              <div className="w-full flex flex-col gap-2">
                {li.map((item, key) => (
                  <ShimmerCard key={key} />
                ))}
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  );
};

export default Services;

// "use client";
// import React, { useEffect, useState } from "react";
// import DescriptionCard from "@/app/components/Shimmer/DescriptionCard";
// import ServicesCard from "@/app/components/Card/ServicesCard";
// import axios from "axios";
// import { useSession } from "next-auth/react";
// import ModalService from "@/app/components/Card/ModalService";

// interface BankService {
//   id: string;
//   name: string;
//   details: string;
//   numberOfUsers: number;
//   status: string;
//   type: string;
//   icon: string;
//   colors: string;
// }

// const Services = () => {
//   const { data: session } = useSession();
//   const li = [1, 2, 3, 4];
//   const colors = [
//     "bg-orange-100",
//     "bg-pink-100",
//     "bg-blue-100",
//     "bg-green-100",
//     "bg-pink-100",
//   ];
//   const [services, setServices] = useState<BankService[]>([]);
//   const [pageNumber, setPageNumber] = useState(1);
//   const [isModalOpen, setIsModalOpen] = useState(false);

//   const handleModalToggle = () => {
//     setIsModalOpen(!isModalOpen);
//   };

//   const accessToken = session?.accessToken as string;
//   console.log(accessToken, "My tokkkkkkkkkkk");
//   async function fetchData(accessToken: string) {
//     try {
//       const response = await axios.get(
//         `https://bank-dashboard-o9tl.onrender.com/bank-services?page=0&size=50`,
//         {
//           headers: {
//             // Authorization: `Bearer ${accessToken}`,
//             Authorization: `Bearer eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJuYXR0eSIsImlhdCI6MTcyNDMwNzUyOCwiZXhwIjoxNzI0MzkzOTI4fQ.kNClFdQpoOYPpOX_aG1RP_DnjwxBMRTlCk8SFQPu04h7dbCYv8BGaJzgvV1AJvR4`,
//           },
//         }
//       );
//       setServices(response.data.data.content);
//       console.log(services);
//     } catch (error) {
//       console.error("There was a problem with the axios request:", error);
//     }
//   }

//   useEffect(() => {
//     fetchData(accessToken);
//   }, [accessToken]);

//   // return (
//   //   <div className="w-[96%] xxs:pt-4 xs:pt-20 md:pt-5 lg:pt-0">
//   //     <div className="ml-5 lg:ml-0 ">
//   //       <div className="mr-5 lg:mr-0 flex gap-4 overflow-x-auto lg:overflow-x-visible [&::-webkit-scrollbar]:hidden [-ms-overflow-style:none] [scrollbar-width:none] lg:pl-10 lg:pt-10 w-full">
//   //         {/* <div className="w-[100%] lg:w-[350px] flex-shrink-0"> */}
//   //         <ServicesCard
//   //           img="/assets/lifeInsurance.svg"
//   //           title="Life Insurance"
//   //           desc="Unlimited Protection"
//   //         />
//   //         {/* </div> */}
//   //         {/* <div className="w-[100%] lg:w-[350px] flex-shrink-0"> */}
//   //         <ServicesCard
//   //           img="/assets/shoppingBag.svg"
//   //           title="Shopping"
//   //           desc="Buy. Think. Grow"
//   //         />
//   //         {/* </div> */}
//   //         {/* <div className="w-[100%] lg:w-[350px] flex-shrink-0"> */}
//   //         <ServicesCard
//   //           img="/assets/safety.svg"
//   //           title="Safety"
//   //           desc="We are your allies"
//   //         />
//   //       </div>
//   //       {/* </div> */}

//   //       <div>
//   //         <div className="flex justify-between">
//   //           <p className="font-semibold text-[22px] text-[#343C6A] dark:text-gray-300 pt-5 pb-5 lg:p-10 ">
//   //             Bank Services List
//   //           </p>
//   //           <div
//   //             className={`flex items-center text-base text-[#718EBF] dark:text-gray-400 rounded-[50px] py-1 pl-6 grow justify-end ${
//   //               isModalOpen ? "blur-sm" : ""
//   //             }`}
//   //           >
//   //             <button
//   //               onClick={handleModalToggle}
//   //               className=" text-white p-3 px-6 rounded-lg shadow-md bg-[#1814F3] hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500"
//   //               type="button"
//   //             >
//   //               Add
//   //             </button>
//   //           </div>
//   //           {isModalOpen && (
//   //             <div
//   //               className="fixed inset-0 z-50 flex justify-center items-center bg-black dark:bg-white dark:bg-opacity-[4%] dark:border-white dark:border-[1px] bg-opacity-50 backdrop-blur-sm"
//   //               onClick={handleModalToggle}
//   //             >
//   //               <div
//   //                 className="relative bg-white dark:bg-[#232328] p-6 rounded-lg shadow-lg max-w-lg w-full"
//   //                 onClick={(e) => e.stopPropagation()} // Prevent modal from closing when clicking inside it
//   //               >
//   //                 <ModalService
//   //                   isOpen={isModalOpen}
//   //                   onClose={handleModalToggle}
//   //                 />
//   //               </div>
//   //             </div>
//   //           )}
//   //         </div>
//   //         <div className="w-full flex flex-col grow items-start px-4">
//   //           {services.length > 0 ? (
//   //             services.map((service, index) => (
//   //               <DescriptionCard
//   //                 key={service.id}
//   //                 img={service.icon}
//   //                 title={service.name}
//   //                 desc={service.details}
//   //                 colOne="Number of Users"
//   //                 descOne={service.numberOfUsers}
//   //                 colTwo="Status"
//   //                 descTwo={service.status}
//   //                 colThree="Type"
//   //                 descThree={service.type}
//   //                 btn="View Details"
//   //                 color={colors[index]}
//   //               />
//   //             ))
//   //           ) : (
//   //             <div className="w-full flex flex-col px-6 gap-2">
//   //               {li.map((item, key) => (
//   //                 <DescriptionCard key={key} />
//   //               ))}
//   //             </div>
//   //           )}
//   //         </div>
//   //       </div>
//   //     </div>
//   //   </div>
//   // );
// };

// export default Services;
