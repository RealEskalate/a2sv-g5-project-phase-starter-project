import 'package:flutter_bloc/flutter_bloc.dart';
import'../injection/injection.dart' as di;
import '../../features/auth/presentation/bloc/auth_bloc.dart';
import '../../features/product/presentation/bloc/product_bloc.dart';

class BlocMultiProvider {
  List<BlocProvider<dynamic>> blocMultiProvider() {
    return [
      BlocProvider<AuthBloc>(
        create: (context) => di.sl<AuthBloc>()),

        BlocProvider<ProductBloc>(
        create: (context) => di.sl<ProductBloc>()),
    ];
  }
}