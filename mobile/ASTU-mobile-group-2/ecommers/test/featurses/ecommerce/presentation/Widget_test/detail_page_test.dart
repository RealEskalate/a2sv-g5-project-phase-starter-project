


import 'dart:io';

import 'package:bloc_test/bloc_test.dart';
import 'package:ecommers/core/Icons/back_icons.dart';
import 'package:ecommers/features/ecommerce/Domain/entity/ecommerce_entity.dart';
import 'package:ecommers/features/ecommerce/presentation/UI/home/Product_detail/detail_page.dart';
import 'package:ecommers/features/ecommerce/presentation/UI/home/header.dart';
import 'package:ecommers/features/ecommerce/presentation/UI/home/home.dart';
import 'package:ecommers/features/ecommerce/presentation/UI/home/product_image.dart';
import 'package:ecommers/features/ecommerce/presentation/state/product_bloc/product_bloc.dart';
import 'package:ecommers/features/ecommerce/presentation/state/product_bloc/product_event.dart';
import 'package:ecommers/features/ecommerce/presentation/state/product_bloc/product_state.dart';
import 'package:ecommers/features/ecommerce/presentation/state/user_states/login_user_states_bloc.dart';
import 'package:ecommers/features/ecommerce/presentation/state/user_states/login_user_states_event.dart';
import 'package:ecommers/features/ecommerce/presentation/state/user_states/login_user_states_state.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';
import 'package:flutter_test/flutter_test.dart';

class MockProductBloc extends MockBloc<ProductEvent, ProductState> implements ProductBloc {}
class MockLoginUserStatesBloc extends MockBloc<LoginUserStatesEvent, LoginUserStates> implements LoginUserStatesBloc {}

void main() {
  late MockProductBloc mockProductBloc;
  late MockLoginUserStatesBloc mockLoginUserStatesBloc;
  


  setUp(() {
    mockProductBloc = MockProductBloc();
    mockLoginUserStatesBloc = MockLoginUserStatesBloc();

    HttpOverrides.global = null;
  });

  tearDown(() {
    mockProductBloc.close();
    mockLoginUserStatesBloc.close();
  });

 

  testWidgets('test add product inistal state ', (WidgetTester tester) async {
    whenListen(
      mockLoginUserStatesBloc,
      Stream.fromIterable([
        LeftUserStates(), // or any other initial state
      ]),
      initialState: LeftUserStates(),
    );
    whenListen(
      mockProductBloc,
      Stream.fromIterable([
        LoadingState(),
        const LoadedAllProductState(
          products: [
            EcommerceEntity(id: '1', name: 'Product 1', description: 'Description 1', imageUrl: 'https://www.simplilearn.com/ice9/free_resources_article_thumb/what_is_image_Processing.jpg', price: 100),
            // EcommerceEntity(id: '2', name: 'Product 2', description: 'Description 2', imageUrl: 'https://www.simplilearn.com/ice9/free_resources_article_thumb/what_is_image_Processing.jpg', price: 200),
          ],
        ),
      ]),
      initialState: ProductIntialState(),
    );
    

    

    // Act
    await tester.pumpWidget(
      MultiBlocProvider(
        providers: [
              BlocProvider<ProductBloc>.value(value: mockProductBloc),
              BlocProvider<LoginUserStatesBloc>.value(
                  value: mockLoginUserStatesBloc),
        
            ],
        child: MaterialApp(
          initialRoute: '/',
          routes: {
            '/detail': (context) => const DetailPage(),
          },
          home: const HomeScreen(),
          builder: EasyLoading.init(),
          
        ),
      ),
    );
    // await tester.tap(find.text('Navigate to AddProduct'));
    await tester.pumpAndSettle();
    await tester.tap(find.byType(ProductImage));
    await tester.pumpAndSettle();
    expect(find.text('Product 1'), findsOneWidget);

    expect(find.byType(BackIcons), findsOneWidget);
    // expect(find.text('Avi'), findsOneWidget);
    // back to home page
    await tester.tap(find.byType(BackIcons));
    await tester.pumpAndSettle();

    expect(find.text('Available Products'), findsOneWidget);
    expect(find.byType(HeaderPart), findsOneWidget);

  
  });
  }