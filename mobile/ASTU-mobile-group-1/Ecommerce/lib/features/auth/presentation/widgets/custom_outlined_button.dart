import 'package:flutter/material.dart';

class CustomOutlinedButton extends StatelessWidget {
  const CustomOutlinedButton(
      {super.key,
      required this.text,
      this.width = 100,
      this.height = 50,
      this.color = Colors.black,
      this.backgroundColor = Colors.white,
      this.onPressed});
  final String text;
  final Color color;
  final double width;
  final double height;
  final Color backgroundColor;
  final void Function()? onPressed;

  @override
  Widget build(BuildContext context) {
    return Row(
      children: [
        Expanded(
          child: Padding(
            padding: const EdgeInsets.all(8.0),
            child: OutlinedButton(
              onPressed: onPressed,
              style: ButtonStyle(
                backgroundColor: WidgetStateProperty.resolveWith(
                  (states) {
                    return backgroundColor;
                  },
                ),
                side: WidgetStateProperty.resolveWith(
                  (states) {
                    return BorderSide(color: color);
                  },
                ),
                foregroundColor: WidgetStateProperty.resolveWith(
                  (states) {
                    return color;
                  },
                ),
                shape:
                    WidgetStateProperty.resolveWith<OutlinedBorder>((states) {
                  return RoundedRectangleBorder(
                    borderRadius: BorderRadius.circular(8),
                  );
                }),
                fixedSize: WidgetStateProperty.resolveWith(
                  (state) {
                    return Size(width, height);
                  },
                ),
              ),
              child: Text(
                text,
              ),
            ),
          ),
        ),
      ],
    );
  }
}
