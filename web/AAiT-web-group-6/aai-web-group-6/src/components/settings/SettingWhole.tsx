import EditProfile from "./EditProfile";

const SettingWhole: React.FC = () => {
  return (
    <>
      <div className="bg-white w-[1120px] rounded-md">
        <div className="ml-14 flex justify-start gap-10 w-[1010px] border-b bg-white pt-5 pb-0 text-slate-400">
          <span className="border-b-2 border-[#1814f3] active-border text-[#1814f3]">
            Edit Profile
          </span>
          <span className="">Preferences</span>
          <span className="">Security</span>
        </div>

        <EditProfile isActive={false} />
      </div>
    </>
  );
};

export default SettingWhole;
