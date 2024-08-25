import 'dart:io';

import 'package:bloc_test/bloc_test.dart';
import 'package:ecommers/features/ecommerce/Domain/entity/ecommerce_entity.dart';
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
import 'package:flutter_test/flutter_test.dart';

// Mocks
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

  testWidgets('HomeScreen displays products correctly', (WidgetTester tester) async {
    // Arrange
    whenListen(
      mockProductBloc,
      Stream.fromIterable([
        LoadingState(),
        const LoadedAllProductState(
          products: [
            EcommerceEntity(id: '1', name: 'Product 1', description: 'Description 1', imageUrl: 'https://www.simplilearn.com/ice9/free_resources_article_thumb/what_is_image_Processing.jpg', price: 100),
            EcommerceEntity(id: '2', name: 'Product 2', description: 'Description 2', imageUrl: 'https://www.simplilearn.com/ice9/free_resources_article_thumb/what_is_image_Processing.jpg', price: 200),
          ],
        ),
      ]),
      initialState: ProductIntialState(),
    );

    whenListen(
      mockLoginUserStatesBloc,
      Stream.fromIterable([
        LeftUserStates(), // or any other initial state
      ]),
      initialState: LeftUserStates(),
    );

    // Act
    await tester.pumpWidget(
      MaterialApp(
        home: MultiBlocProvider(
          providers: [
            BlocProvider<ProductBloc>.value(value: mockProductBloc),
            BlocProvider<LoginUserStatesBloc>.value(value: mockLoginUserStatesBloc),
          ],
          child: const HomeScreen(),
        ),
      ),
    );

   // Process the second state (LoadedAllProductState)
    await tester.pumpAndSettle();
    // Assert
    expect(find.text('Available Products'), findsOneWidget);
    expect(find.text('Product 1'), findsOneWidget);
    expect(find.text('Product 2'), findsOneWidget);
    expect(find.byType(ProductImage), findsNWidgets(2));
    expect(find.byType(Image), findsNWidgets(4));
    expect(find.text('(4.0)'), findsNWidgets(2));
    expect(find.text('\$100.0'), findsOneWidget);
    expect(find.text('\$200.0'), findsOneWidget);
    expect(find.byType(Icon), findsWidgets);
    expect(find.byType(ListView), findsOneWidget);
    expect(find.byType(ClipRRect), findsOneWidget);
  });



   testWidgets('HomeScreen displays error state ', (WidgetTester tester) async {
    // Arrange
    whenListen(
      mockProductBloc,
      Stream.fromIterable([
        LoadingState(),
        const ProductErrorState(messages: 'try again')
      ]),
      initialState: ProductIntialState(),
    );

    whenListen(
      mockLoginUserStatesBloc,
      Stream.fromIterable([
        LeftUserStates(), // or any other initial state
      ]),
      initialState: LeftUserStates(),
    );

    // Act
    await tester.pumpWidget(
      MaterialApp(
        home: MultiBlocProvider(
          providers: [
            BlocProvider<ProductBloc>.value(value: mockProductBloc),
            BlocProvider<LoginUserStatesBloc>.value(value: mockLoginUserStatesBloc),
          ],
          child: const HomeScreen(),
        ),
      ),
    );

   // Process the second state (LoadedAllProductState)
    await tester.pumpAndSettle();

    expect(find.byType(SnackBar), findsOneWidget);
    expect(find.byType(ElevatedButton), findsOneWidget);
    expect(find.text('try again'), findsNWidgets(2));
   });




   testWidgets('HomeScreen must showe loading state', (WidgetTester tester) async {
    // Arrange
    whenListen(
      mockProductBloc,
      Stream.fromIterable([
        LoadingState(),
        
      ]),
      initialState: ProductIntialState(),
    );

    whenListen(
      mockLoginUserStatesBloc,
      Stream.fromIterable([
        LeftUserStates(), // or any other initial state
      ]),
      initialState: LeftUserStates(),
    );

    // Act
    await tester.pumpWidget(
      MaterialApp(
        home: MultiBlocProvider(
          providers: [
            BlocProvider<ProductBloc>.value(value: mockProductBloc),
            BlocProvider<LoginUserStatesBloc>.value(value: mockLoginUserStatesBloc),
          ],
          child: const HomeScreen(),
        ),
      ),
    );

   // Process the second state (LoadedAllProductState)
    await tester.pumpAndSettle();

    expect(find.byKey(const Key('loading')), findsOneWidget);
    
   
   });

  
  
  
}
