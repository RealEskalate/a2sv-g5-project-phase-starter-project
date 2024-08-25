import 'package:flutter/material.dart';

class ChatAppBar extends StatelessWidget implements PreferredSizeWidget {
  final String name;
  final String status;
  final String imageUrl;
  final bool isOnline;


  ChatAppBar(this.name, this.status, this.imageUrl, {this.isOnline = false});

  @override
  Widget build(BuildContext context) {
    return AppBar(
      backgroundColor: Colors.white,
      elevation: 0.5,
      leading: IconButton(
        icon: Icon(Icons.arrow_back),
        onPressed: () {
          Navigator.pop(context);
          // Navigator.pushReplacementNamed(context, '/HomeChat');
        },
      ),
      title: Row(
        mainAxisAlignment: MainAxisAlignment.start,
        children: [
          Stack(
            children: [
              Transform.translate(
                offset: Offset(-10, 0),
                child: CircleAvatar(radius: 25, backgroundColor: Colors.pink, backgroundImage: NetworkImage(imageUrl))),
                Positioned(
                  bottom: 3,
                  right: 10,
                  child: Container(
                    height: 12,
                    width: 12,
                    decoration: BoxDecoration(
                      color: isOnline ? Colors.green : Colors.grey,
                      borderRadius: BorderRadius.circular(50),
                      border: Border.all(color: Colors.white, width: 1)
                    ),
                  ),
                )

            ],
          ),
          
          Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(
                name,
                style: TextStyle(
                  fontSize: 18,
                  color: Colors.black,
                  fontWeight: FontWeight.w500,
                ),
              ),
              Text(
                status,
                style: TextStyle(
                  fontSize: 12,
                  color: Colors.grey,
                ),
              ),
            ],
          ),
        ],
      ),
      actions: [
        IconButton(icon: Icon(Icons.call), onPressed: () {}),
        IconButton(icon: Icon(Icons.videocam), onPressed: () {}),
      ],
    );
  }

  @override
  Size get preferredSize => Size.fromHeight(kToolbarHeight);
}
