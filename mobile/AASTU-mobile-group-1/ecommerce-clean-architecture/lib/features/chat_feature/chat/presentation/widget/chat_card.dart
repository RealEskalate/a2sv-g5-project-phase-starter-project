
import 'package:flutter/material.dart';

class ChatCard extends StatelessWidget {
  final String name;
  final String message;
  final String time;
  final String imagePath;
  final VoidCallback onDelete;

  ChatCard({
    required this.name,
    required this.message,
    required this.time,
    required this.imagePath,
    required this.onDelete,
  });



  @override
  Widget build(BuildContext context) {
    return Dismissible(
      key: UniqueKey(), // UniqueKey to identify the widget
      direction: DismissDirection.endToStart, // Slide from right to left
      onDismissed: (direction) {
        // Call the onDelete callback when the card is dismissed
        onDelete();
      },
      background: Container(
        color: Colors.red,
        alignment: Alignment.centerRight,
        padding: EdgeInsets.only(right: 20),
        child: Icon(
          Icons.delete,
          color: Colors.white,
        ),
      ),
      child: Container(
        width: 348,
        height: 60.65,
        decoration: BoxDecoration(
          color: Colors.transparent,
          border: Border.all(
            color: Colors.white,
            width: 0.98,
          ),
        ),
        child: Row(
          children: [
            CircleAvatar(
              radius: 30,
              backgroundImage: AssetImage(imagePath),
            ),
            SizedBox(width: 10),
            Expanded(
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  Expanded(
                    child: Column(
                      crossAxisAlignment: CrossAxisAlignment.start,
                      children: [
                        Row(
                          children: [
                            Expanded(
                              child: Text(
                                name,
                                style: TextStyle(
                                  color: Colors.black,
                                  fontSize: 17,
                                  fontWeight: FontWeight.w400,
                                ),
                              ),
                            ),
                            Text(
                              time,
                              style: TextStyle(
                                color: Color(0xFF797C7B),
                                fontSize: 12,
                              ),
                              textAlign: TextAlign.right,
                            ),
                          ],
                        ),
                        SizedBox(height: 2),
                        Text(
                          message,
                          style: TextStyle(
                            color: Color(0xFF797C7B),
                            fontSize: 12,
                          ),
                          overflow: TextOverflow.ellipsis,
                        ),
                      ],
                    ),
                  ),
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }
}
