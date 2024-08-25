import '../../../../core/constants/colors.dart';
import '../../../../core/constants/profile_photo.dart';
import '../../../auth/presentation/pages/pages.dart';

Widget showRecievedMessage(String text){
  return Column(
    crossAxisAlignment: CrossAxisAlignment.center,
    children: [
      Row(
        mainAxisAlignment: MainAxisAlignment.start,
        children: [
          CircleAvatar(
            backgroundColor: ChatColors.profileColors[1],
            child: Image(image: AssetImage(ChatUserProfilePhoto.photos[1])),
          ),
          const Text('Annaie Ellson'),
          ],
      ),
      Container(
        padding: const EdgeInsets.all(10),
        decoration: BoxDecoration(
          color: ChatColors.recievedColor,
          borderRadius: const  BorderRadius.only(topLeft: Radius.circular(10),bottomLeft: Radius.circular(10),bottomRight: Radius.circular(10))
        ),
        child: Text(text),
      ),
      const Text('03:25 AM')
    ],
  );
}