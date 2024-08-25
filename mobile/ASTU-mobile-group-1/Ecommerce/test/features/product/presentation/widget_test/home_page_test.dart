import 'dart:io';

import 'package:bloc_test/bloc_test.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:product_6/core/cubit/user_cubit.dart';
import 'package:product_6/features/auth/domain/entities/user_entity.dart';
import 'package:product_6/features/product/domain/entities/product_entity.dart';
import 'package:product_6/features/product/presentation/bloc/product_bloc.dart';
import 'package:product_6/features/product/presentation/pages/home_page.dart';

class MockProductBloc extends MockBloc<ProductEvent, ProductState>
    implements ProductBloc {}

class MockUserCubit extends MockBloc<UserCubit, UserEntity?>
    implements UserCubit {}

void main() {
  late MockProductBloc mockProductBloc;
  late MockUserCubit mockUserCubit;

  setUp(() {
    mockProductBloc = MockProductBloc();
    mockUserCubit = MockUserCubit();
    HttpOverrides.global = null;
  });

  tearDown(() {
    mockProductBloc.close();
    mockUserCubit.close();
  });

  testWidgets(
    'Home page display product correctly',
    (WidgetTester tester) async {
      // Arrange
      whenListen(
        mockProductBloc,
        Stream.fromIterable([
          LoadingState(),
          const LoadedAllProductsState(
            products: [
              ProductEntity(
                  id: '1',
                  name: 'Product 1',
                  description: 'Description 1',
                  imageUrl:
                      'https://www.simplilearn.com/ice9/free_resources_article_thumb/what_is_image_Processing.jpg',
                  price: 100,
                  seller: UserEntity.empty),
              ProductEntity(
                  id: '2',
                  name: 'Product 2',
                  description: 'Description 2',
                  imageUrl:
                      'https://www.simplilearn.com/ice9/free_resources_article_thumb/what_is_image_Processing.jpg',
                  price: 200,
                  seller: UserEntity.empty),
            ],
          ),
        ]),
        initialState: InitalState(),
      );

      whenListen(
        mockUserCubit,
        Stream.fromIterable([
          const UserEntity(
            id: 'user1',
            name: 'Test User',
            email: 'test@example.com',
          )
        ]),
        initialState: null,
      );

      await tester.pumpWidget(
        MaterialApp(
          home: MultiBlocProvider(
            providers: [
              BlocProvider<ProductBloc>.value(value: mockProductBloc),
              BlocProvider<UserCubit>.value(value: mockUserCubit),
            ],
            child: const HomePage(),
          ),
        ),
      );

      await tester.pumpAndSettle();

      expect(find.text('Available Products'), findsOneWidget);
      expect(find.text('Product 1'), findsOneWidget);
      expect(find.text('Product 2'), findsOneWidget);
      expect(find.byType(Image), findsNWidgets(2));
      expect(find.text('(4.0)'), findsNWidgets(2));
      expect(find.text('\$100.0'), findsOneWidget);
      expect(find.text('\$200.0'), findsOneWidget);
      expect(find.byType(Icon), findsWidgets);
      expect(find.byType(ListView), findsOneWidget);
      expect(find.byType(ClipRRect), findsNWidgets(2));
    },
  );

  testWidgets('Home page displays error state ', (WidgetTester tester) async {
    // Arrange
    whenListen(
      mockProductBloc,
      Stream.fromIterable(
          [LoadingState(), const ErrorState(message: 'try again')]),
      initialState: InitalState(),
    );
    whenListen(
      mockUserCubit,
      Stream.fromIterable([
        const UserEntity(
          id: 'user1',
          name: 'Test User',
          email: 'test@example.com',
        )
      ]),
      initialState: null,
    );
    // Act
    await tester.pumpWidget(
      MaterialApp(
        home: MultiBlocProvider(
          providers: [
            BlocProvider<ProductBloc>.value(value: mockProductBloc),
            BlocProvider<UserCubit>.value(value: mockUserCubit),
          ],
          child: const HomePage(),
        ),
      ),
    );

    await tester.pumpAndSettle();

    expect(find.text('Failed to load products. Please try again later.'),
        findsOneWidget);
  });
}
