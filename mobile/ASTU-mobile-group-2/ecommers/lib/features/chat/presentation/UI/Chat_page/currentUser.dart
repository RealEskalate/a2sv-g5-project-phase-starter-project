import 'package:flutter/material.dart';

class CurrentUser extends StatelessWidget {
  final String name;
  final String image;
  final bool online;
  final Color? statusColor;

  const CurrentUser({
    Key? key,
    required this.name,
    required this.image,
    required this.online,
    this.statusColor,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    final width = MediaQuery.of(context).size.width;
 
    return Row(
      children: [
        Stack(
          children: [
            CircleAvatar(
              radius: 20.0,
              backgroundImage: AssetImage(image),
            ),
            Positioned(
              bottom: 0,
              right: 0,
              child: Container(
                width: 12,
                height: 12,
                decoration: BoxDecoration(
                  color: online ? Colors.green : Colors.red,
                  shape: BoxShape.circle,
                  border: Border.all(
                    color: Colors.white,
                    width: 1,
                  ),
                ),
              ),
            ),
          ],
        ),
        const SizedBox(width: 8),
        Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Container(
             
              width: width * 0.4,
              child:  Text(
                name,
                overflow: TextOverflow.ellipsis,
                style: const  TextStyle(
                  fontSize: 16,
                  fontWeight: FontWeight.bold,
                  color: Colors.black,
                ),
              ),
            ),
            Text(
              online ? 'Online' : 'Offline',
              style: TextStyle(
                fontSize: 12,
                color: statusColor ?? Colors.black54,
              ),
            ),
          ],
        ),
      ],
    );
  }
}
