

import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';


// import 'features/product/domain/usecase/add_product.dart';
// import 'features/product/domain/usecase/get_product.dart';
// import 'features/product/data/data_sources/remote_data_source.dart';
// import 'features/product/data/models/product_model.dart';
import 'features/authentication/presentation/bloc/blocs.dart';
// import 'features/authentication/presentation/pages/onboarding.dart';
import 'features/authentication/presentation/pages/onboarding.dart';
import 'features/chat/presentation/pages/chat_page.dart';
import 'features/chat/presentation/widgets/user_avater.dart';
import 'features/product/presentation/bloc/blocs.dart';
import 'service_locator.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
   
  await setUp();
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context)  {
    var userBloc = getIt<UserBloc>();
    return MultiBlocProvider( 
      providers: [
        BlocProvider(create: (context) => getIt<ProductBloc>()),
        BlocProvider(create: (context) => userBloc),
      ],
      child: const MaterialApp(
        debugShowCheckedModeBanner: false,
        home: ChatPage()   
      ),
    );
  }
}
