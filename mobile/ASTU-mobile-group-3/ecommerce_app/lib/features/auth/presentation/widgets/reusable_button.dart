import 'package:flutter/material.dart';

class ReusableButton extends StatelessWidget {
  final String lable;
  const ReusableButton({super.key, required this.lable});

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.only(bottom: 18.0, top: 28),
      child: DecoratedBox(
        decoration: BoxDecoration(
            color: const Color(0xff3F51F3),
            borderRadius: BorderRadius.circular(10)),
        child: Padding(
          padding: const EdgeInsets.symmetric(vertical: 16, horizontal: 34),
          child: Center(
              child: reusableText(lable, FontWeight.w600, 14, Colors.white)),
        ),
      ),
    );
  }
}

Text reusableText(String text, FontWeight wight, double size,
    [Color color = Colors.black]) {
  return Text(
    text,
    overflow: TextOverflow.clip,
    style: TextStyle(fontWeight: wight, fontSize: size, color: color),
  );
}
