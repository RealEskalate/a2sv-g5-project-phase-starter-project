

import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../../../core/Colors/colors.dart';
import '../../../../../core/border/border_style.dart';
import '../../state/input_button_activation/button_bloc.dart';
import '../../state/input_button_activation/button_event.dart';

class InputField extends StatelessWidget {
  final String text;
  final String disc;
  final String placeHolder;
  final Map<String,dynamic> data;
  const InputField({
    super.key,
    required this.text,
    this.disc = '',
    required this.placeHolder,
    required this.data
    });

  @override
  Widget build(BuildContext context) {
    return  Expanded(
      child: TextFormField(
        onChanged: (value) => {
          context.read<ButtonBloc>().add(InsertInput(name: data['name'],price: data['price'].toString(),description: data['disc'],type: data['type'],id: data['id'])),
          if (text == 'name') {
            context.read<ButtonBloc>().add(InsertInput(name: value,tag: 'name',type: data['type']))
          } else if (text == 'price') {
            context.read<ButtonBloc>().add(InsertInput(price: value,tag: 'price',type: data['type']))
          } else if (text == 'description') {
            context.read<ButtonBloc>().add(InsertInput(description: value,tag: 'description',type: data['type']))
          }
       
          },
        scribbleEnabled: true,
        scrollPhysics: const BouncingScrollPhysics(),
        maxLines: null,
       
        initialValue: placeHolder,
        decoration: InputDecoration(
          
          hintText: disc,
          focusedBorder:text != 'description'? borderStyle:largInputBOrderStyle,
          enabledBorder: text != 'description'? borderStyle:largInputBOrderStyle,
          disabledBorder: text != 'description'? borderStyle:largInputBOrderStyle,
          hintStyle: const TextStyle(
            
            color: smallText,
            fontSize: 14
          )
        ),
      ),
    );
  }
}