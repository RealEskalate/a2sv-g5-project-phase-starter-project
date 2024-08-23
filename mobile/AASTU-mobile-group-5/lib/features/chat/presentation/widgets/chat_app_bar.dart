import 'package:flutter/material.dart';
import 'profile_widget.dart';

class CustomAppBar extends StatelessWidget implements PreferredSizeWidget {
  @override
  Size get preferredSize => const Size.fromHeight(80); // Adjust the height if needed

  @override
  Widget build(BuildContext context) {
    return AppBar(
      backgroundColor: Colors.white, 
      elevation: 0, // Removes shadow
      leading: Padding(
        padding: const EdgeInsets.only(left: 16.0, top: 21.0),
        child: CircleAvatar(
          radius: 30.0, 
          backgroundColor: Colors.grey[300], 
          child: ProfileWidget(iconUrl: 'assets/images/Alex.png',isOnline: true,),
        ),
      ),
      title: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Padding(
            padding: const EdgeInsets.only(top: 13.0), 
            child: const Text(
              'Sabila Sayma',
              style: TextStyle(
                fontFamily: 'General Sans Variable', 
                fontSize: 18,
                fontWeight: FontWeight.w500,
                height: 20 / 18, 
                color: Colors.black, 
              ),
            ),
          ),
          Padding(
            padding: const EdgeInsets.only(top: 8.0), // Align the second text below
            child: const Text(
              'Last seen recently',
              style: TextStyle(
                fontFamily: 'General Sans Variable',
                fontSize: 12,
                fontWeight: FontWeight.w400,
                height: 12 / 12,
                color: Colors.black54,
              ),
            ),
          ),
        ],
      ),
      actions: [
        IconButton(
          icon: Icon(Icons.call_outlined, color: Colors.black),
          onPressed: () {
            // Handle call action(if we add  any)
          },
        ),
        IconButton(
          icon: Icon(Icons.videocam_outlined, color: Colors.black),
          onPressed: () {
            // Handle video call action(if we add any)
          },
        ),
        const SizedBox(width: 16), 
      ],
    );
  }
}
