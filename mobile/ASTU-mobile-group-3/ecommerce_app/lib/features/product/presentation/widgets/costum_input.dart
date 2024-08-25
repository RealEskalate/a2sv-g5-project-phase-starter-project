import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../../core/themes/themes.dart';
import '../bloc/cubit/input_validation_cubit.dart';

class CostumInput extends StatelessWidget {
  final String hint;
  final String text;
  final TextEditingController control;
  final Color? fillColor;
  final Color? borderColor;
  final int maxLine;
  final Color textColor;
  final bool obsecure;
  const CostumInput(
      {super.key,
      required this.hint,
      required this.control,
      required this.text,
      this.obsecure = false,
      this.fillColor,
      this.borderColor,
      this.maxLine = 1,
      this.textColor = MyTheme.ecBlack});

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Padding(
          padding: const EdgeInsets.symmetric(vertical: 10, horizontal: 20),
          child: Text(
            text,
            style: TextStyle(fontWeight: FontWeight.bold, color: textColor),
          ),
        ),
        Padding(
          padding: const EdgeInsets.symmetric(horizontal: 20),
          child: BlocBuilder<InputValidationCubit, InputValidationState>(
            builder: (context, state) {
              late Color? color;
              if (borderColor == null) {
                color = Colors.transparent;
              } else {
                color = borderColor!;
              }
              if (state is InputValidatedState) {
                if (!state.getState(text)[0]) {
                  color = MyTheme.ecRed;
                }
              }
              return TextField(
                key: Key(text),
                controller: control,
                maxLines: maxLine,
                obscureText: obsecure,
                onChanged: (value) {
                  BlocProvider.of<InputValidationCubit>(context)
                      .checkChanges([text, value]);
                },
                decoration: InputDecoration(
                    contentPadding: const EdgeInsets.all(10),
                    fillColor:
                        (fillColor == null) ? MyTheme.ecInputGrey : fillColor,
                    filled: true,
                    border: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(10),
                    ),
                    enabledBorder: OutlineInputBorder(
                      borderSide: BorderSide(
                        color: (borderColor == null)
                            ? Colors.transparent
                            : borderColor!,
                      ),
                    ),
                    focusedBorder: OutlineInputBorder(
                      borderSide: BorderSide(
                        color: color,
                      ),
                    ),
                    hintText: hint,
                    hintStyle: const TextStyle(color: MyTheme.ecGrey)),
              );
            },
          ),
        ),
      ],
    );
  }
}
