
import '../../../../core/constants/colors.dart';
import '../../../../core/constants/profile_photo.dart';
import '../../../auth/presentation/pages/pages.dart';

Widget showUser(){
  return CircleAvatar(
    minRadius: 30,
    maxRadius: 40,
    backgroundColor: ChatColors.getRandomColor(),
  child: Image(
    image: AssetImage(ChatUserProfilePhoto.getRandomPhoto())),
  );
}