import Image from "next/image";

const Cardinfo = () => {
    return (
      <div>
        <div className="flex p-5 bg-white rounded-xl mb-2 mt-4 max-w-screen-sm justify-between min-w-[325px]">
          <div className="flex-initial w-[2/12] m-3">
            <div className="text-blue-500 bg-opacity-25 font-semibold py-1 px-2 rounded-lg  text-sm w-[45px]">
              <Image
                src={`/icons/Cardbill.svg`}
                alt={"Cards"}
                width={27}
                height={18}
              />
            </div>
          </div>
          <div className="flex-initial w-[4/12] m-3">
            <div>
              <h2>Card Type</h2>
              <p className="text-gray-500">Secondary</p>
            </div>
          </div>
          <div className="flex-initial w-[3/12] m-3">
            <div>
              <h2>Bank</h2>
              <p className="text-gray-500">DBL Bank</p>
            </div>
          </div>
          <div className="flex-initial w-[3/12] m-3">
            <p className="text-[#1814F3]"> view details</p>
          </div>
        </div>
      </div>
    );
  };

export default Cardinfo