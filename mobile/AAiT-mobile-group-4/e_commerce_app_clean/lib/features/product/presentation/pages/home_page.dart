import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../authentication/presentation/bloc/auth_bloc.dart';
import '../bloc/product_bloc.dart';
import '../widgets/components/header.dart';
import '../widgets/components/product_card.dart';
import '../widgets/components/styles/snack_bar_style.dart';

class Home extends StatefulWidget {
  const Home({super.key});
  @override
  State<Home> createState() => _HomeState();
}

class _HomeState extends State<Home> {
  @override
  Widget build(BuildContext context) {
    context.read<ProductBloc>().add(LoadAllProductEvent());
    context.read<AuthBloc>().add(GetCurrentUserEvent());
    return Scaffold(
      body: Center(
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            const Padding(
              padding: EdgeInsets.fromLTRB(12, 32, 12, 0),
              child: HeaderView(),
            ),
            const SizedBox(height: 22.0),
            Expanded(
              child: BlocListener<AuthBloc, AuthState>(
                listener: (context, state) {
                 if (state is AuthErrorState) {
                    ScaffoldMessenger.of(context)
                        .showSnackBar(customSnackBar(state.message,  Theme.of(context).secondaryHeaderColor));
                  }
                },
                child: BlocBuilder<ProductBloc, ProductState>(
                  builder: (context, state) {
                    if (state is ProductLoading) {
                      return const Center(
                        child: CircularProgressIndicator(),
                      );
                    }
                    if (state is LoadedAllProductState) {
                      return RefreshIndicator(
                        onRefresh: () async {
                          context
                              .read<ProductBloc>()
                              .add(LoadAllProductEvent());
                        },
                        child: ListView.builder(
                          itemCount: state.products.length,
                          itemBuilder: (context, index) {
                            
                            return MyCardBox(product: state.products[index]);
                          },
                        ),
                      );
                    } else if (state is ProductErrorState) {
                      return Center(
                        child: Text(state.message),
                      );
                    } else {
                      return Container();
                    }
                  },
                ),
              ),
            ),
          ],
        ),
      ),
      floatingActionButton: SizedBox(
        width: 64,
        height: 64,
        child: FloatingActionButton(
          //-----------------------------------------------------------------
          onPressed: () {
            Navigator.pushNamed(context, '/product_add_page');
          },
          //-----------------------------------------------------------------
          backgroundColor: Theme.of(context).primaryColor,
          shape: const CircleBorder(),
          child: const Icon(
            Icons.add,
            size: 36,
            color: Colors.white,
          ),
        ),
      ),
    );
  }
}
