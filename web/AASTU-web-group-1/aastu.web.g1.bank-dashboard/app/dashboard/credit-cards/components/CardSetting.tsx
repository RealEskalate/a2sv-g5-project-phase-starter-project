import Image from "next/image";

interface props {
    image: string;
    color: string;
    title: string;
    description: string;
  }
  
  const CardSetting = ({ image, color, title, description }: props) => {
    return (
      <div className="flex bg-white mb-3 rounded-xl">
        <div className="flex-initial w-[5/12] m-3 text-[16px]">
          <div
            className={`${color} bg-opacity-25 font-semibold py-1 px-2 rounded-lg  text-sm w-[45px] `}
          >
            <Image
              src={image}
              alt={title}
              width={20}
              height={20}
              className="mx-auto"
            />
          </div>
        </div>
        <div className="flex-initial w-[7/12] m-3">
          <div>
            <h1>{title}</h1>
            <p className="text-[#718EBF]">{description}</p>
          </div>
        </div>
      </div>
    );
  };
  
  export default CardSetting