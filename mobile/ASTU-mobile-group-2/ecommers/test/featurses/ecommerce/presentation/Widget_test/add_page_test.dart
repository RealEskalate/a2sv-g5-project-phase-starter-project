import 'dart:io';

import 'package:bloc_test/bloc_test.dart';
import 'package:ecommers/features/ecommerce/Domain/entity/ecommerce_entity.dart';
import 'package:ecommers/features/ecommerce/presentation/UI/add_product/add_product.dart';
import 'package:ecommers/features/ecommerce/presentation/UI/add_product/image_field.dart';
import 'package:ecommers/features/ecommerce/presentation/UI/add_product/input_border.dart';
import 'package:ecommers/features/ecommerce/presentation/UI/home/home.dart';
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
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';
import 'package:flutter_test/flutter_test.dart';

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
          },
          home: const HomeScreen(),
          builder: EasyLoading.init(),
          
        ),
      ),
    );

    // Press the button to navigate
    await tester.tap(find.byKey(const Key('add Product page')));
    await tester.pumpAndSettle();

    // Assert
    expect(find.byType(Scaffold), findsOneWidget);
    expect(find.byType(IinputBorder), findsNWidgets(4));
    expect(find.byType(ImageField), findsOneWidget);
    // expect(find.text('Product Name: Test Product'), findsOneWidget);
  });
  



  testWidgets('add product error state in screen', (WidgetTester tester) async {
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
        ProductIntialState(), // or any other initial state
      ]),
      initialState: ProductIntialState(),
    );
    whenListen(
      mockButtonBloc,
      Stream.fromIterable([
        AddErrorState(messages: 'error to add'), // or any other initial state
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
          },
          home: Builder(
              builder: (context) {
                return ElevatedButton(
                  onPressed: () {
                    Navigator.pushNamed(context, '/add-product',
                    arguments: {'id':'','imageUrl':'','price':0,'name':'','disc':'','type':0},);
                  },
                  child: const Text('Navigate to AddProduct'),
                );
              },
            ),
            builder: EasyLoading.init(),
          
        ),
      ),
    );

    // Press the button to navigate
    await tester.tap(find.text('Navigate to AddProduct'));
    await tester.pumpAndSettle();
    

    // Assert
    expect(find.text('try again'), findsOneWidget);
    expect(find.byType(SnackBar), findsOneWidget);

 
    
    // expect(find.text('Product Name: Test Product'), findsOneWidget);
  });

  
}
