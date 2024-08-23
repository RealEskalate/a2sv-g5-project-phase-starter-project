import 'package:flutter/material.dart';

class ProfileWidget extends StatelessWidget {
  final bool isOnline;
  final String iconUrl;
  final bool isInAppBar;
  final bool curStatus;
  final bool isFirst;

  const ProfileWidget({
    super.key,
    required this.isOnline,
    required this.iconUrl,
    this.isInAppBar = true,
    this.curStatus = true,
    this.isFirst = false,
  });

  @override
  Widget build(BuildContext context) {
    double avatarRadius = isInAppBar ? 25 : 30;

    return Stack(
      children: [
        Container(
          padding: const EdgeInsets.all(2),
          decoration: BoxDecoration(
            shape: BoxShape.circle,
            border: curStatus
                ? Border.all(
                    color: Color.fromARGB(255, 4, 52, 92),
                    width: 2,
                  )
                : null,
          ),
          child: CircleAvatar(
            radius: avatarRadius,
            backgroundColor: Colors.blue.shade300,
            child: CircleAvatar(
              radius: avatarRadius - 10, // Adjust inner radius accordingly
              backgroundImage: Image.asset(
                iconUrl,
              ).image,
            ),
          ),
        ),
        if (isOnline)
          Positioned(
            bottom: 5,
            right: 5,
            child: Container(
              width: 15,
              height: 10,
              decoration: BoxDecoration(
                shape: BoxShape.circle,
                color: Colors.green,
                
              ),
            ),
          ),
        if (isFirst)
          Positioned(
            bottom: 0,
            right: 0,
            child: Stack(
              alignment: Alignment.center,
              children: [
                Container(
                  width: 20,
                  height:
                      20, // Adjust the height to match the width for a perfect circle
                  decoration: BoxDecoration(
                    shape: BoxShape.circle,
                    color: Colors.white, // Set the color to white
                    border: Border.all(
                      color: Colors.grey,
                      width: 2,
                    ),
                  ),
                ),
                Positioned(
                  child: InkWell(
                    onTap: () {
                      // Add your onPressed code here!
                    },
                    child: Text(
                      '+',
                      style: TextStyle(
                        color: Colors.grey,
                        fontSize: 15,
                      ),
                    ),
                  ),
                  
                ),
              ],
            ),
          ),
      ],
    );
  }
}
