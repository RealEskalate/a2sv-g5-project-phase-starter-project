import Image from "next/image";
const ProfileAvatar = () => {
  return (
    <div className="mr-10 flex justify-center md:block py-3 h-auto">
      <Image
        alt="Profile Image"
        src={"/assets/images/profile-avatar.png"}
        width={`170`}
        height={`170`}
      />
    </div>
  );
};

export default ProfileAvatar;