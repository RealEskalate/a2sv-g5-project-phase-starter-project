import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../../core/constants/constants.dart';
import '../../../../core/themes/themes.dart';
import '../../../../core/validator/validator.dart';
import '../bloc/cubit/user_input_validation_cubit.dart';

class CostumInput extends StatelessWidget {
  final String hint;
  final String text;
  final TextEditingController control;
  final Color? fillColor;
  final Color? borderColor;
  final int maxLine;
  final Color textColor;
  final String fromWhere;
  final bool obscure;
  const CostumInput(
      {super.key,
      required this.hint,
      required this.control,
      required this.text,
      required this.fromWhere,
      this.obscure = false,
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
          child:
              BlocBuilder<UserInputValidationCubit, UserInputValidationState>(
            builder: (context, state) {
              late Color? color;
              if (borderColor == null) {
                color = Colors.transparent;
              } else {
                color = borderColor!;
              }
              if (fromWhere == AppData.signup &&
                  state is SignupUserInputValidated) {
                if (!state.getSingleInputState(text)) {
                  color = MyTheme.ecRed;
                }
              }
              if (fromWhere == AppData.login &&
                  state is LoginUserInputValidated) {
                if (!state.getSingleInputState(text)) {
                  color = MyTheme.ecRed;
                }
              }

              return TextField(
                key: Key(text),
                controller: control,
                maxLines: maxLine,
                obscureText: obscure,
                onChanged: (value) {
                  if (text == InputDataValidator.confirmPass) {
                    BlocProvider.of<UserInputValidationCubit>(context)
                        .checkWith(fromWhere, text, value.trim());
                  } else {
                    BlocProvider.of<UserInputValidationCubit>(context)
                        .checkWith(fromWhere, text, value.trim());
                  }
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
