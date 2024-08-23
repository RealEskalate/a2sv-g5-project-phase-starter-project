import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'theme_cubit.dart';

class DarkMode extends StatelessWidget {
  const DarkMode({super.key});

  @override
  Widget build(BuildContext context) {
    return BlocBuilder<ThemeCubit, ThemeMode>(
      builder: (context, themeMode) {
        return IconButton(
          icon: Icon(
            themeMode == ThemeMode.dark ? Icons.nights_stay : Icons.wb_sunny,
          ),
          onPressed: () {
            context.read<ThemeCubit>().toggleTheme();
          },
        );
      },
    );
  }
}