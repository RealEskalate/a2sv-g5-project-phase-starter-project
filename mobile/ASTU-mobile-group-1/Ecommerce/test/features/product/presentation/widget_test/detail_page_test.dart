import 'dart:io';

import 'package:bloc_test/bloc_test.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:product_6/config/route/route.dart' as route;
import 'package:product_6/core/cubit/user_cubit.dart';
import 'package:product_6/features/auth/domain/entities/user_entity.dart';
import 'package:product_6/features/product/domain/entities/product_entity.dart';
import 'package:product_6/features/product/presentation/bloc/product_bloc.dart';
import 'package:product_6/features/product/presentation/pages/home_page.dart';
import 'package:product_6/features/product/presentation/widgets/back_button.dart';
import 'package:product_6/features/product/presentation/widgets/product_list_builder.dart';

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
    'Detail page displays product correctly after loading',
    (WidgetTester tester) async {
      // Arrange
      const testProduct = ProductEntity(
          id: '66c8e1b9d9557bd7c3fb0211',
          name: 'Product 1',
          description: 'Description 1',
          price: 100000.0,
          imageUrl:
              'https://res.cloudinary.com/g5-mobile-track/image/upload/v1724441017/images/nxlukefakavym716bsle.webp',
          seller: UserEntity.empty);

      // Set up mock state transitions
      whenListen(
        mockProductBloc,
        Stream.fromIterable([
          LoadingState(),
          const LoadedAllProductsState(
            products: [testProduct],
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
          ),
        ]),
        initialState: null,
      );

      await tester.pumpWidget(
        MultiBlocProvider(
          providers: [
            BlocProvider<ProductBloc>.value(value: mockProductBloc),
            BlocProvider<UserCubit>.value(value: mockUserCubit),
          ],
          child: const MaterialApp(
            home: HomePage(),
            onGenerateRoute: route.controller,
          ),
        ),
      );

      await tester.pumpAndSettle();

      // Verify that the product list is displayed on the HomePage
      expect(find.text('Available Products'), findsOneWidget);
      expect(find.text('Product 1'), findsOneWidget);
      expect(find.byType(ProductListItem), findsWidgets);

      await tester.tap(find.byType(ProductListItem).first);

      // Transition to the detail page
      whenListen(
        mockProductBloc,
        Stream.fromIterable([
          const LoadedSingleProductState(product: testProduct),
        ]),
        initialState: LoadingState(),
      );

      await tester.pumpAndSettle();

      expect(find.text('Product 1'), findsOneWidget);
      expect(find.text('Description 1'), findsOneWidget);
      expect(find.text('\$100000.0'), findsOneWidget);

      // do the back button and check the product in homepage
      expect(find.byType(BackButtonWidget), findsOneWidget);
      await tester.tap(find.byType(BackButtonWidget));

      await tester.pumpAndSettle();

      expect(find.text('Available Products'), findsOneWidget);
      expect(find.text('Product 1'), findsOneWidget);
    },
  );
}
