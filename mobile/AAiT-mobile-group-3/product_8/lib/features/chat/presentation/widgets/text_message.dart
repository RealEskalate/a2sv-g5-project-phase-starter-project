import 'package:flutter/material.dart';

class TextMessage extends StatelessWidget {
  final bool isLeft;
  const TextMessage({
    super.key,
    required this.isLeft,
  });

  @override
  Widget build(BuildContext context) {
    return Card(
      margin:
          EdgeInsets.only(top: 18, left: isLeft ? 9 : 0, right: isLeft ? 0 : 9),
      shadowColor: const Color(0xFFF2F7FB).withOpacity(0.5),
      clipBehavior: Clip.antiAlias,
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.only(
          bottomLeft: const Radius.circular(15),
          bottomRight: const Radius.circular(15),
          topRight:
              isLeft ? const Radius.circular(15) : const Radius.circular(0),
          topLeft:
              isLeft ? const Radius.circular(0) : const Radius.circular(15),
        ),
      ),
      color: isLeft
          ? const Color.fromARGB(255, 220, 239, 254)
          : const Color(0xFF3F51F3),
      child: Padding(
        padding: const EdgeInsets.all(8.0),
        child: Text(
          'Have a great working week!!padding: const EdgeInsets.only(top: 18, left: 9)',
          softWrap: true,
          overflow: TextOverflow.ellipsis,
          maxLines: 10,
          style: TextStyle(
            color: isLeft ? Colors.black : Colors.white,
            fontWeight: FontWeight.w500,
            fontSize: 12,
          ),
        ),
      ),
    );
  }
}
