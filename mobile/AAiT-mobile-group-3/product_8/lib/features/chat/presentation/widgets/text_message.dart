import 'package:flutter/material.dart';

class TextMessage extends StatelessWidget {
  const TextMessage({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Card(
      margin: const EdgeInsets.only(top: 18, left: 9),
      shadowColor: const Color(0xFFF2F7FB).withOpacity(0.5),
      clipBehavior: Clip.antiAlias,
      shape: const RoundedRectangleBorder(
          borderRadius: BorderRadius.only(
        bottomLeft: Radius.circular(15),
        bottomRight: Radius.circular(15),
        topRight: Radius.circular(15),
      )),
      color: const Color.fromARGB(255, 220, 239, 254),
      // color: Colors.black,
      // elevation: 3,
      child: const Padding(
        padding: const EdgeInsets.all(8.0),
        child: const Text(
          softWrap: true,
          overflow: TextOverflow.ellipsis,
          'Have a great working week!!padding: const EdgeInsets.only(top: 18, left: 9)',
          maxLines: 10,
        ),
      ),
    );
  }
}
