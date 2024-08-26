


import 'dart:io';

import 'package:bloc_test/bloc_test.dart';
import 'package:ecommers/core/Icons/back_icons.dart';
import 'package:ecommers/features/ecommerce/Domain/entity/ecommerce_entity.dart';
import 'package:ecommers/features/ecommerce/presentation/UI/add_product/add_product.dart';
import 'package:ecommers/features/ecommerce/presentation/UI/add_product/input_border.dart';
import 'package:ecommers/features/ecommerce/presentation/UI/home/Product_detail/detail_page.dart';
import 'package:ecommers/features/ecommerce/presentation/UI/home/header.dart';
import 'package:ecommers/features/ecommerce/presentation/UI/home/home.dart';
import 'package:ecommers/features/ecommerce/presentation/UI/home/product_image.dart';
import 'package:ecommers/features/ecommerce/presentation/UI/seachProduct/search_screen.dart';
import 'package:ecommers/features/ecommerce/presentation/state/image_input_display/image_bloc.dart';
import 'package:ecommers/features/ecommerce/presentation/state/image_input_display/image_event.dart';
import 'package:ecommers/features/ecommerce/presentation/state/image_input_display/image_state.dart';
import 'package:ecommers/features/ecommerce/presentation/state/input_button_activation/bottum_state.dart';
import 'package:ecommers/features/ecommerce/presentation/state/input_button_activation/button_bloc.dart';
import 'package:ecommers/features/ecommerce/presentation/state/input_button_activation/button_event.dart';
import 'package:ecommers/features/ecommerce/presentation/state/product_bloc/product_bloc.dart';
import 'package:ecommers/features/ecommerce/presentation/state/product_bloc/product_event.dart';
import 'package:ecommers/features/ecommerce/presentation/state/product_bloc/product_state.dart';
import 'package:ecommers/features/ecommerce/presentation/state/user_states/login_user_states_bloc.dart';
import 'package:ecommers/features/ecommerce/presentation/state/user_states/login_user_states_event.dart';
import 'package:ecommers/features/ecommerce/presentation/state/user_states/login_user_states_state.dart';
import 'package:ecommers/features/login/presentation/UI/login_home_page.dart';
import 'package:ecommers/features/login/presentation/UI/register_body.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:integration_test/integration_test.dart';


class MockProductBloc extends MockBloc<ProductEvent, ProductState>
    implements ProductBloc {}

class MockLoginUserStatesBloc
    extends MockBloc<LoginUserStatesEvent, LoginUserStates>
    implements LoginUserStatesBloc {}

class MockButtonBloc extends MockBloc<ButtonEvent, BottumState>
    implements ButtonBloc {}

class MockImageBloc extends MockBloc<ImageEvent, ImageState>
    implements ImageBloc {}

void main() {
 
  IntegrationTestWidgetsFlutterBinding.ensureInitialized();


  late MockProductBloc mockProductBloc;
  late MockLoginUserStatesBloc mockLoginUserStatesBloc;
  late MockButtonBloc mockButtonBloc;
  late MockImageBloc mockImageBloc;
  

  setUp(() {
    mockProductBloc = MockProductBloc();
    mockLoginUserStatesBloc = MockLoginUserStatesBloc();
    mockButtonBloc = MockButtonBloc();
    mockImageBloc = MockImageBloc();
    HttpOverrides.global = null;
  });

  tearDown(() {
    mockProductBloc.close();
    mockLoginUserStatesBloc.close();
    mockButtonBloc.close();
    mockImageBloc.close();
  });
  final String image = '/home/samuel/Documents/2024-internship-mobile-tasks-/Mobile/Samuel_Tolossa/ecommers/assets/image/splashScreen.png';

 

  testWidgets('test add product state ', (WidgetTester tester) async {
    whenListen(
      mockLoginUserStatesBloc,
      Stream.fromIterable([
        LeftUserStates(), // or any other initial state
      ]),
      initialState: LeftUserStates(),
    );
    whenListen(
      mockButtonBloc,
      Stream.fromIterable([
        IntialState(), // or any other initial state
      ]),
      initialState: IntialState(),
    );
    whenListen(
      mockImageBloc,
      Stream.fromIterable([
        InputIntialState(), // or any other initial state
      ]),
      initialState: InputIntialState(),
    );
    whenListen(
      mockProductBloc,
      Stream.fromIterable([
        LoadingState(),
        const LoadedAllProductState(
          products: [
            EcommerceEntity(id: '1', name: 'Product 1', description: 'Description 1', imageUrl: 'https://www.simplilearn.com/ice9/free_resources_article_thumb/what_is_image_Processing.jpg', price: 100, sellerName: ''),
            EcommerceEntity(id: '2', name: 'Product 2', description: 'Description 2', imageUrl: 'https://www.simplilearn.com/ice9/free_resources_article_thumb/what_is_image_Processing.jpg', price: 200),
            EcommerceEntity(id: '3', name: 'Product 3', description: 'Description 3', imageUrl: 'https://www.simplilearn.com/ice9/free_resources_article_thumb/what_is_image_Processing.jpg', price: 300),
            EcommerceEntity(id: '4', name: 'Product 4', description: 'Description 4', imageUrl: 'https://www.simplilearn.com/ice9/free_resources_article_thumb/what_is_image_Processing.jpg', price: 400),
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
              BlocProvider<ButtonBloc>.value(value: mockButtonBloc),
              BlocProvider<ImageBloc>.value(value: mockImageBloc),
        
            ],
        child: MaterialApp(
          initialRoute: '/',
          routes: {
            '/add-product': (context) => const AddProduct(),
            '/detail': (context) => const DetailPage(),
            '/search': (context) => const SearchScreen(),
            '/home': (context) => const HomeScreen(),
            '/login': (context) => const LoginHomePage(),
            '/registration': (context) => const RegisterBody(),
          },
          home: const HomeScreen(),
          builder: EasyLoading.init(),
          
        ),
      ),
    );
    // await tester.tap(find.text('Navigate to AddProduct'));
    await tester.pumpAndSettle();
    await Future.delayed(const Duration(seconds: 2));
    await tester.tap(find.byType(ProductImage).at(0));
    await tester.pumpAndSettle();
    await Future.delayed(const Duration(seconds: 2));
    expect(find.text('Product 1'), findsOneWidget);

    expect(find.byType(BackIcons), findsOneWidget);
    // expect(find.text('Avi'), findsOneWidget);
    // back to home page
    await Future.delayed(const Duration(seconds: 2));
    await tester.tap(find.byType(BackIcons));
    await tester.pumpAndSettle();
    await Future.delayed(const Duration(seconds: 2));
    expect(find.text('Available Products'), findsOneWidget);
    await tester.dragUntilVisible(find.byType(ListView), find.byType(ProductImage).at(1), const Offset(0, -1000));
    await tester.pumpAndSettle();
    await tester.dragUntilVisible(find.byType(ListView), find.byType(ProductImage).at(2), const Offset(0, -1000));
    await tester.pumpAndSettle();
    await tester.dragUntilVisible(find.byType(ListView), find.byType(ProductImage).at(3), const Offset(0, -1000));
    await tester.pumpAndSettle();
  
    await Future.delayed(const Duration(seconds: 2));
    expect(find.byType(HeaderPart), findsOneWidget);
    await Future.delayed(const Duration(seconds: 2));
    await tester.tap(find.byKey(const Key('add Product page')));
    await tester.pumpAndSettle();
    await Future.delayed(const Duration(seconds: 2));
    expect(find.text('name'), findsOneWidget);
    await tester.enterText(find.byType(IinputBorder).at(0), 'Smart Phone');
    await Future.delayed(const Duration(seconds: 2));
    await tester.enterText(find.byType(IinputBorder).at(1), 'Phone');
    await Future.delayed(const Duration(seconds: 2));
    await tester.enterText(find.byType(IinputBorder).at(2), '2345.4');
    await Future.delayed(const Duration(seconds: 2));
    await tester.enterText(find.byType(IinputBorder).at(3), 'smart phone for every one');
    await Future.delayed(const Duration(seconds: 2));
    whenListen(
      mockImageBloc,
      Stream.fromIterable([
        OnImageSelect(image: image, file: File(image)),
         // or any other initial state
      ]),
      initialState: InputIntialState(),
    );
    await tester.pumpAndSettle();
    
    await tester.tap(find.byKey(const Key('image selector')));
    await Future.delayed(const Duration(seconds: 4));
    await tester.pumpAndSettle();
    await tester.tap(find.byKey(const Key('back from add page')));
    await tester.pumpAndSettle();
    await Future.delayed(const Duration(seconds: 2));
    // redirectToSerch
    await tester.tap(find.byKey(const Key('redirectToSerch')));
    await tester.pumpAndSettle();
    await Future.delayed(const Duration(seconds: 2));
    expect(find.text('Search'), findsOneWidget);
    await Future.delayed(const Duration(seconds: 2));



  
  });
  }